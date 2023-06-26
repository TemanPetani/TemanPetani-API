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

	mapData := NewProductModel(data)
	if tx := repo.db.Create(&mapData); tx.Error != nil {
		return "", tx.Error
	}

	return data.ID, nil
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

// Select implements products.ProductDataInterface
func (*ProductData) Select() (products []products.Core, err error) {
	panic("unimplemented")
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
