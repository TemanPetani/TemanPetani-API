package products

import "mime/multipart"

type Core struct {
	ID					string
	Name				string			`json:"name" form:"name" validate:"required"`
	Price 			float64			`json:"price" form:"price" validate:"required"`
	Stock				uint				`json:"stock" form:"stock" validate:"required,gt=0"`
	Description	string			`json:"description" form:"description"`
	ImageUrl		string			`json:"imageUrl" form:"imageUrl"`
	UserID			uint			
}

type CoreProductImage struct {
	Image 			*multipart.FileHeader				
}

type ProductDataInterface interface {
	Insert(data Core) (productId string, err error)
	Select() (products []Core, err error)
	SelectById(productId string) (product Core, err error) 
	Update(productId string, data Core) error
	Delete(productId string) error
	InsertImage(productId string, image CoreProductImage) (ImageUrl string, err error)
	DeleteImage(productId string) error
}

type ProductServiceInterface interface {
	AddProduct(data Core) (productId string, err error)
}
