package service

import (
	"alta/temanpetani/features/products"
	"errors"

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
