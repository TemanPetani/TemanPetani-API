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
	validator   *validator.Validate
}

// DeleteProductById implements products.ProductServiceInterface
func (service *ProductService) DeleteProductById(productId string, userId uint) error {
	if verifyOwner := service.productData.VerifyProductOwner(productId, userId); verifyOwner {
		if err := service.productData.Delete(productId); err != nil {
			return err
		}
	}
	return nil
}

// UpdateProductById implements products.ProductServiceInterface
func (service *ProductService) UpdateProductById(productId string, data products.Core) error {
	if errValidator := service.validator.Struct(data); errValidator != nil {
		return errors.New("error validation: " + errValidator.Error())
	}
	if verifyOwner := service.productData.VerifyProductOwner(productId, data.UserID); verifyOwner {
		if err := service.productData.Update(productId, data); err != nil {
			return err
		}
	}
	return nil
}

// GetProductById implements products.ProductServiceInterface
func (service *ProductService) GetProductById(productId string) (*products.Core, error) {
	product, err := service.productData.SelectById(productId)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// GetAllProducts implements products.ProductServiceInterface
func (service *ProductService) GetAllProducts(querys map[string]any) ([]products.Core, error) {
	allProducts, err := service.productData.Select(querys)
	if err != nil {
		return nil, err
	}
	return allProducts, nil
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

// UpdateImage implements products.ProductServiceInterface
func (service *ProductService) UpdateImage(productId string, image products.CoreProductImageRequest) (imageUrl string, err error) {
	if errValidator := service.validator.Struct(image); errValidator != nil {
		return "", errors.New("error validator: " + errValidator.Error())
	}
	newImage, errGetImage := image.Image.Open()
	if errGetImage != nil {
		return "", errors.New("failed to open file: " + errGetImage.Error())
	}
	defer newImage.Close()
	imageKey := helpers.GenerateNewId() + "_" + image.Image.Filename
	_, errUpload := helpers.UploaderS3().PutObject(&s3.PutObjectInput{
		Bucket: aws.String("alta-airbnb"),
		Key:    aws.String(imageKey),
		Body:   newImage,
	})
	if errUpload != nil {
		return "", errors.New("failed to upload file image: " + errUpload.Error())
	}
	image.ImageUrl = "https://alta-airbnb.s3.ap-southeast-3.amazonaws.com/" + imageKey
	errInsert := service.productData.UpdateImage(productId, image)
	if errInsert != nil {
		return "", errInsert
	}

	return image.ImageUrl, nil
}

func New(productData products.ProductDataInterface) products.ProductServiceInterface {
	return &ProductService{
		productData: productData,
		validator:   validator.New(),
	}
}
