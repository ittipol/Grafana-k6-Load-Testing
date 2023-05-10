package services

import "go-load-testing/repositories"

type ProductService interface {
	GetProduct() error
}

type productService struct {
	productRepository repositories.Product
}

func NewProductService(productRepository repositories.Product) ProductService {
	return &productService{productRepository}
}

func (obj productService) GetProduct() error {

	obj.productRepository.GetProduct()

	return nil
}
