package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var key = "products"

type productRepositoryRedis struct {
	// Product
	db          *gorm.DB
	redisClient *redis.Client
}

func NewProductRepositoryRedis(db *gorm.DB, redisClient *redis.Client) Product {
	db.AutoMigrate(&product{})
	mockData(db)
	return &productRepositoryRedis{db, redisClient}
}

func (obj productRepositoryRedis) GetCachedProduct() (products []product, err error) {
	cachedData, err := obj.redisClient.Get(context.Background(), key).Result()

	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(cachedData), &products)

	return
}

func (obj productRepositoryRedis) GetProduct() (products []product, err error) {

	cachedData, err := obj.redisClient.Get(context.Background(), key).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		return
	} else {
		err = json.Unmarshal([]byte(cachedData), &products)

		if err != nil {
			return
		}

		fmt.Println("Get Products from Redis")

		return
	}

	err = obj.db.Order("quantity DESC").Limit(50).Find(&products).Error

	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(products)

	if err != nil {
		return
	}

	fmt.Println("Set Products to Redis")

	err = obj.redisClient.Set(context.Background(), key, bytes, time.Minute*5).Err()
	if err != nil {
		panic(err)
	}

	return
}
