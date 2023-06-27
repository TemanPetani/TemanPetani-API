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
	data.ID = helpers.GenerateNewId()

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
		if querys["role"] != nil {
			if querys["role"] == "admin" && productMap.User.Role == "user" {
				allProductsMap = append(allProductsMap, productMap)
			} else if querys["role"] == "user" && productMap.User.Role == "admin" {
				allProductsMap = append(allProductsMap, productMap)
			}
		} else if querys["owner"] != nil {
			if querys["owner"] == "admin" && productMap.User.Role == "admin" {
				allProductsMap = append(allProductsMap, productMap)
			} else if productMap.User.ID == querys["userId"] {
				allProductsMap = append(allProductsMap, productMap)
			}
		}
	}
	return allProductsMap, nil
}

// SelectById implements products.ProductDataInterface
func (repo *ProductData) SelectById(productId string) (*products.Core, error) {
	var product Products
	if tx := repo.db.Where("id = ?", productId).First(&product); tx.Error != nil {
		return nil, tx.Error
	}
	productMap := ModelToProductCore(product)
	return &productMap, nil
}

// Update implements products.ProductDataInterface
func (repo *ProductData) Update(productId string, data products.Core) error {
	dataMap := CoreToProductModel(data)
	if tx := repo.db.Model(&Products{}).Where("id = ?", productId).Updates(dataMap); tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Delete implements products.ProductDataInterface
func (repo *ProductData) Delete(productId string) error {
	if tx := repo.db.Where("id = ?", productId).Delete(&Products{}); tx.Error != nil {
		return tx.Error
	}
	return nil
}

// VerifyProductOwner implements products.ProductDataInterface
func (repo *ProductData) VerifyProductOwner(productId string, owner uint) bool {
	var products Products
	if tx := repo.db.Where("id = ? && user_id = ?", productId, owner).First(&products); tx.RowsAffected != 0 {
		return true
	}
	return false
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

func New(db *gorm.DB) products.ProductDataInterface {
	return &ProductData{
		db: db,
	}
}
