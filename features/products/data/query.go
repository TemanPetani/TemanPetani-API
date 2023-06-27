package data

import (
	"alta/temanpetani/features/products"
	"alta/temanpetani/utils/helpers"

	"gorm.io/gorm"
)

type ProductData struct {
	db *gorm.DB
}

// Insert implements products.ProductDataInterface
func (repo *ProductData) Insert(data products.Core) (productId string, err error) {
	data.ID = helpers.GenerateNewId();

	mapData := CoreToProductModel(data)
	if tx := repo.db.Create(&mapData); tx.Error != nil {
		return "", tx.Error
	}

	return data.ID, nil
}

// Select implements products.ProductDataInterface
func (repo *ProductData) Select(querys map[string]any) ([]products.Core, error) {
	var allProducts []Products
	if tx := repo.db.Preload("User").Find(&allProducts); tx.Error != nil {
		return nil, tx.Error
	}
	var allProductsMap []products.Core
	for _, product := range allProducts {
		productMap := ModelToProductCore(product)
		if querys["role"] == "admin" && productMap.User.Role == "user" {
			allProductsMap = append(allProductsMap, productMap)
		} else if querys["role"] == "user" && productMap.User.Role == "admin" {
			allProductsMap = append(allProductsMap, productMap)
		}
	}
	return allProductsMap, nil
}

// Delete implements products.ProductDataInterface
func (*ProductData) Delete(productId string) error {
	panic("unimplemented")
}

// DeleteImage implements products.ProductDataInterface
func (*ProductData) DeleteImage(productId string) error {
	panic("unimplemented")
}

// InsertImage implements products.ProductDataInterface
func (repo *ProductData) UpdateImage(productId string, image products.CoreProductImageRequest) error {
	if tx := repo.db.Model(&Products{}).Where("id = ?", productId).Update("image_url", image.ImageUrl); tx.Error != nil {
		return tx.Error
	}
	return nil
}

// SelectById implements products.ProductDataInterface
func (*ProductData) SelectById(productId string) (product products.Core, err error) {
	panic("unimplemented")
}

// Update implements products.ProductDataInterface
func (*ProductData) Update(productId string, data products.Core) error {
	panic("unimplemented")
}

func New(db *gorm.DB) products.ProductDataInterface {
	return &ProductData{
		db: db,
	}
}
