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
	Role			string			`gorm:"type:enum('admin','user');default:'user'"`
}

func CoreToProductModel(data products.Core) Products {
	return Products{
		ID: data.ID,
		Name: data.Name,
		Price: data.Price,
		Stock: data.Stock,
		Description: data.Description,
		ImageUrl: data.ImageUrl,
		UserID: data.UserID,
	}
}

func ModelToProductCore(model Products) products.Core {
	return products.Core{
		ID: model.ID,
		Name: model.Name,
		Price: model.Price,
		Stock: model.Stock,
		Description: model.Description,
		ImageUrl: model.ImageUrl,
		User: products.Users(model.User),		
	}
}