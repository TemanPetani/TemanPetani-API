package data

import (
	"alta/temanpetani/features/products"
	"time"

	"gorm.io/gorm"
)

type Products struct {
	ID					string				`gorm:"type:varchar(50);primaryKey"`
	Name				string				`gorm:"type:varchar(150);notNull"`
	Price 			float64				`gorm:"type:decimal(10,2);notNull"`
	Stock				uint					`gorm:"type:uint;notNull"`
	Description	string				`gorm:"type:text"`
	ImageUrl		string				`gorm:"type:varchar(150)"`
	UserID			uint 					`gorm:"type:uint"`
	User				Users					`gorm:"foreignKey:UserID"`
	CreatedAt 	time.Time	
	UpdatedAt		time.Time
	DeletedAt		gorm.DeletedAt `gorm:"index"`
}

type Users struct {
	ID				string			`gorm:"type:uint;primaryKey"`
	FullName	string 			`gorm:"type:varchar(100);notNull"`
	Email			string 			`gorm:"type:varchar(50);unique:notNull"`
}

func NewProductModel(data products.Core) Products {
	return Products{
		ID: data.ID,
		Name: data.Name,
		Price: data.Price,
		Stock: data.Stock,
		Description: data.Description,
		ImageUrl: data.ImageUrl,
	}
}