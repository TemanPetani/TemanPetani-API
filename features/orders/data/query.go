package data

import (
	"alta/temanpetani/features/orders"
	"alta/temanpetani/utils/helpers"

	"gorm.io/gorm"
)

type OrderData struct {
	db *gorm.DB
}

// Insert implements orders.OrderDataInterface.
func (repo *OrderData) Insert(productId string, data orders.Core) (orderId string, err error) {
	data.ID = helpers.GenerateNewId()

	mapData := CoreToOrderModel(data)
	if tx := repo.db.Create(&mapData); tx.Error != nil {
		return "", tx.Error
	}

	return data.ID, nil
}

// Select implements orders.OrderDataInterface.
func (repo *OrderData) Select() ([]orders.Core, error) {
	var allOrders []Orders
	if tx := repo.db.Preload("User").Preload("Product").Find(&allOrders); tx.Error != nil {
		return nil, tx.Error
	}
	var allOrdersMap []orders.Core 
	
}

func New(db *gorm.DB) orders.OrderDataInterface {
	return &OrderData{
		db: db,
	}
}
