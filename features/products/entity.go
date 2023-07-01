package products

import "mime/multipart"

type Core struct {
	ID          string                `json:"id" form:"id"`
	Name        string                `json:"name" form:"name" validate:"required"`
	Price       float64               `json:"price" form:"price" validate:"required"`
	Stock       uint                  `json:"stock" form:"stock" validate:"required,gt=0"`
	Description string                `json:"description" form:"description"`
	ImageUrl    string                `json:"imageUrl" form:"imageUrl"`
	Image       *multipart.FileHeader `json:"image,omitempty" form:"image,omitempty"`
	UserID      uint64                `json:"userId,omitempty"`
	User        Users                 `json:"owner" form:"owner"`
}

type Users struct {
	ID       uint64 `json:"id" form:"id"`
	FullName string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Role     string `json:"role" form:"role"`
}

type CoreProductImageRequest struct {
	Image    *multipart.FileHeader `form:"image" validate:"required"`
	ImageUrl string
}

type ProductDataInterface interface {
	Insert(data Core) (productId string, err error)
	Select(querys map[string]any) ([]Core, error)
	SelectById(productId string) (product Core, err error)
	SelectByUserId(userId uint64) ([]Core, error)
	Update(productId string, data Core) error
	Delete(productId string) error
	UpdateImage(productId string, image CoreProductImageRequest) error
	DeleteImage(productId string) error
}

type ProductServiceInterface interface {
	AddProduct(data Core) (productId string, err error)
	UpdateImage(productId string, image CoreProductImageRequest) (imageUrl string, err error)
	GetAllProducts(querys map[string]any) ([]Core, error)
	GetProductsByUserId(userId uint64) ([]Core, error)
}
