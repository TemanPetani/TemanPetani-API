package products

import "mime/multipart"

type Core struct {
	ID					string
	Name				string			`json:"name" form:"name" validate:"required"`
	Price 			float64			`json:"price" form:"price" validate:"required"`
	Stock				uint				`json:"stock" form:"stock" validate:"required,gt=0"`
	Description	string			`json:"description" form:"description"`
	ImageUrl		string			`json:"imageUrl" form:"imageUrl"`
	Image				*multipart.FileHeader			`form:"image"`
	UserID			uint			
}

type CoreProductImageRequest struct {
	Image 			*multipart.FileHeader				`form:"image" validate:"required"`
	ImageUrl 		string			
}

type ProductDataInterface interface {
	Insert(data Core) (productId string, err error)
	Select() (products []Core, err error)
	SelectById(productId string) (product Core, err error) 
	Update(productId string, data Core) error
	Delete(productId string) error
	UpdateImage(productId string, image CoreProductImageRequest) error
	DeleteImage(productId string) error
}

type ProductServiceInterface interface {
	AddProduct(data Core) (productId string, err error)
	UpdateImage(productId string, image CoreProductImageRequest) (imageUrl string, err error)
}
