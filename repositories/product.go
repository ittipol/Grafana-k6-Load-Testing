package repositories

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type product struct {
	ID       int
	Name     string `gorm:"size:100"`
	Quantity int
}

type Product interface {
	GetProduct() ([]product, error)
	GetCachedProduct() (products []product, err error)
}

func mockData(db *gorm.DB) error {

	var count int64
	tx := db.Model(&product{}).Count(&count)

	if tx.Error != nil {
		return tx.Error
	}

	if count > 0 {
		return nil
	}

	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	products := []product{}
	for i := 0; i < 500; i++ {
		products = append(products, product{
			Name:     fmt.Sprintf("Product %v", i+1),
			Quantity: random.Intn(100),
		})
	}

	return db.Create(&products).Error
}
