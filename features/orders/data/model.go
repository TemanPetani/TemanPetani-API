package data

import "alta/temanpetani/features/orders"

type Orders struct {
	ID        string   `gorm:"primaryKey"`
	UserID    uint     `gorm:"type:uint"`
	User      Users    `gorm:"foreignKey:UserID"`
	ProductID string   `gorm:"type:varchar(50)"`
	Product   Products `gorm:"foreignKey:ProductID"`
	Bank      string   `gorm:"type:varchar(50);notNull"`
	Quantity  uint     `gorm:"type:uint;notNull"`
	Note      string   `gorm:"type:text"`
	Status    string   `gorm:"type:varchar(50);default:'pending'"`
}

type Users struct {
	ID       string `gorm:"type:uint;primaryKey"`
	FullName string `gorm:"type:varchar(100);notNull"`
}

type Products struct {
	ID    string  `gorm:"type:varchar(50);primaryKey"`
	Name  string  `gorm:"type:varchar(150);notNull"`
	Price float64 `gorm:"type:decimal(10,2);notNull"`
}

func CoreToOrderModel(data orders.Core) Orders {
	return Orders{
		ID: data.ID,
		UserID: data.UserID,
		ProductID: data.Product.ID,
		Bank: data.Bank,
		Quantity: data.Quantity,
		Note: data.Note,
		Status: data.Status,
	}
}

func ModelToOrderCore(model Orders) orders.Core {
	return orders.Core{
		ID: model.ID,
		User: orders.Users(model.User),
	}
}