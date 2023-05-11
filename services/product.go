package services

type productResponse struct {
	ID       int
	Name     string
	Quantity int
}

type ProductService interface {
	GetProducts() (products []productResponse, err error)
}
