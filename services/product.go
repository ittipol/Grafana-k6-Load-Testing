package services

type productResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type ProductService interface {
	GetProducts() (products []productResponse, err error)
}
