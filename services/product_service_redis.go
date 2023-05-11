package services

import (
	"context"
	"encoding/json"
	"fmt"
	"go-load-testing/repositories"
	"time"

	"github.com/redis/go-redis/v9"
)

type productServiceRedis struct {
	productRepository repositories.Product
	redisClient       *redis.Client
}

func NewProductServiceRedis(productRepository repositories.Product, redisClient *redis.Client) ProductService {
	return &productServiceRedis{productRepository, redisClient}
}

func (obj productServiceRedis) GetProducts() (products []productResponse, err error) {

	key := "service:GetProducts"

	if cachedData, err := obj.redisClient.Get(context.Background(), key).Result(); err == nil {

		if err = json.Unmarshal([]byte(cachedData), &products); err == nil {
			fmt.Println("Get Products data from Redis")
			return products, err
		}

	} else if err != nil && err != redis.Nil {
		return nil, err
	}

	// Get Product data from database
	data, err := obj.productRepository.GetProducts()

	if err != nil {
		return nil, err
	}

	for _, product := range data {
		products = append(products, productResponse{
			ID:       product.ID,
			Name:     product.Name,
			Quantity: product.Quantity,
		})
	}

	if bytes, err := json.Marshal(products); err == nil {
		fmt.Println("Set Products data to Redis")

		obj.redisClient.Set(context.Background(), key, bytes, time.Second*30)
	}

	// if err = obj.redisClient.Set(context.Background(), key, bytes, time.Minute*2).Err(); err != nil {
	// 	return nil, err
	// }

	return products, nil
}
