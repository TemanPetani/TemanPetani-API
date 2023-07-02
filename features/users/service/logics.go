package service

import (
	"alta/temanpetani/features/users"
	"alta/temanpetani/utils/helpers"
	"errors"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-playground/validator/v10"
)

type userService struct {
	userData users.UserDataInterface
	validate *validator.Validate
}

func New(repo users.UserDataInterface) users.UserServiceInterface {
	return &userService{
		userData: repo,
		validate: validator.New(),
	}
}

func (service *userService) Login(email string, password string) (users.UserCore, string, error) {
	if email == "" || password == "" {
		return users.UserCore{}, "", errors.New("Email dan Password Harus Diisi")
	}
	dataLogin, token, err := service.userData.Login(email, password)
	return dataLogin, token, err
}

func (service *userService) Create(input users.UserCore) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}

	errInsert := service.userData.Insert(input)
	if errInsert != nil {
		return errInsert
	}

	return nil
}

func (service *userService) GetById(id uint64) (users.UserCore, error) {
	data, err := service.userData.SelectById(id)
	if err != nil {
		return users.UserCore{}, err
	}
	return data, err
}

func (service *userService) UpdateById(id uint64, input users.UserCore) error {
	errUpdate := service.userData.UpdateById(id, input)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}

func (service *userService) DeleteById(id uint64) error {
	errUpdate := service.userData.DeleteById(id)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}

func (service *userService) UpdateImage(id uint64, image *multipart.FileHeader) error {
	newImage, errGetImage := image.Open()
	if errGetImage != nil {
		return errors.New(errGetImage.Error())
	}
	defer newImage.Close()

	imageKey := helpers.GenerateNewId() + "_" + image.Filename
	_, errUpload := helpers.UploaderS3().PutObject(&s3.PutObjectInput{
		Bucket: aws.String("alta-airbnb"),
		Key:    aws.String(imageKey),
		Body:   newImage,
	})

	if errUpload != nil {
		return errors.New(errUpload.Error())
	}

	imageUrl := "https://alta-airbnb.s3.ap-southeast-3.amazonaws.com/" + imageKey
	errInsert := service.userData.UpdateImage(id, imageUrl)
	if errInsert != nil {
		return errInsert
	}

	return nil
}
