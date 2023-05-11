package services

import (
	"go-load-testing/repositories"
)

type productService struct {
	productRepository repositories.Product
}

func NewProductService(productRepository repositories.Product) ProductService {
	return &productService{productRepository}
}

func (obj productService) GetProducts() (products []productResponse, err error) {

	data, err := obj.productRepository.GetProducts()

	if err != nil {
		return
	}

	for _, product := range data {
		products = append(products, productResponse{
			ID:       product.ID,
			Name:     product.Name,
			Quantity: product.Quantity,
		})
	}

	return
}
