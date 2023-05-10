package services

import (
	"fmt"
	"go-load-testing/repositories"

	"github.com/redis/go-redis/v9"
)

type productResponse struct {
	ID       int
	Name     string
	Quantity int
}

type ProductService interface {
	GetProduct() (products []productResponse, err error)
}

type productService struct {
	productRepository repositories.Product
}

func NewProductService(productRepository repositories.Product) ProductService {
	return &productService{productRepository}
}

func (obj productService) GetProduct() (products []productResponse, err error) {
	fmt.Printf("Call Service\n")
	cachedData, err := obj.productRepository.GetCachedProduct()

	if err != nil {
		if err != redis.Nil {
			return
		}
	} else {
		// Redis Key Found

		for _, product := range cachedData {
			products = append(products, productResponse{
				ID:       product.ID,
				Name:     product.Name,
				Quantity: product.Quantity,
			})
		}

		return
	}

	// Get From DB

	// Set To Redis

	return
}
