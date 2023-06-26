package service

import (
	"alta/temanpetani/features/products"
	"alta/temanpetani/utils/helpers"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-playground/validator/v10"
)

type ProductService struct {
	productData products.ProductDataInterface
	validator *validator.Validate
}

// AddProduct implements products.ProductServiceInterface
func (service *ProductService) AddProduct(data products.Core) (productId string, err error) {
	if errValidator := service.validator.Struct(data); errValidator != nil {
		return "", errors.New("error validator: " + errValidator.Error())
	}
	image, errGetImage := data.Image.Open()
	if errGetImage != nil {
		return "", errors.New("failed to open file: " + errGetImage.Error())
	}
	defer image.Close()
	imageKey := helpers.GenerateNewId() + "_" + data.Image.Filename
	_, errUpload := helpers.UploaderS3().PutObject(&s3.PutObjectInput{
		Bucket: aws.String("alta-airbnb"),
		Key:    aws.String(imageKey),
		Body:   image,
	})
	if errUpload != nil {
		return "", errors.New("failed to upload file image: " + errUpload.Error())
	}
	data.ImageUrl = "https://alta-airbnb.s3.ap-southeast-3.amazonaws.com/" + imageKey
	id, errInsert := service.productData.Insert(data)
	if errInsert != nil {
		return "", errInsert
	}

	return id, nil
}

func New(productData products.ProductDataInterface) products.ProductServiceInterface {
	return &ProductService{
		productData: productData,
		validator: validator.New(),
	}
}
