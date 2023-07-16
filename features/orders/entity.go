package orders

type Core struct {
	ID       string   `json:"id" form:"id"`
	UserID   uint     `json:"userId" form:"userId"`
	User     Users    `json:"user" form:"user"`
	Product  Products `json:"product" form:"product"`
	Bank     string   `json:"bank" form:"bank" validate:"required"`
	Quantity uint     `json:"quantity" form:"bank" validate:"required"`
	Note     string   `json:"note,omitempty" form:"note,omitempty"`
	Status   string   `json:"status" form:"status"`
}

type Users struct {
	ID       uint   `json:"id" form:"id"`
	FullName string `json:"fullname" form:"fullname"`
}

type Products struct {
	ID    string `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Price string `json:"price" form:"price"`
}

type OrderDataInterface interface {
	Insert(productId string, data Core) (orderId string, err error)
	Select() ([]Core, error)
}

type OrderServiceInterface interface {
	AddOrder(productId string, data Core) (orderId string, err error)
	GetAllOrders() ([]Core, error)
}