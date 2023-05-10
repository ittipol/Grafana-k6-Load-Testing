package repositories

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type productRepositoryRedis struct {
	// Product
	db *redis.Client
}

func NewProductRepositoryRedis(db *redis.Client) Product {
	return &productRepositoryRedis{db}
}

func (obj productRepositoryRedis) GetProduct() ([]product, error) {
	val, err := obj.db.Get(context.Background(), "foo").Result()
	if err != nil {
		return nil, err
	}
	fmt.Println("foo", val)

	return nil, nil
}
