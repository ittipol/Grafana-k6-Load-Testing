package repositories

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type productRepositoryDB struct {
	// Product
	db *gorm.DB
}

func NewProductRepositoryDB(db *gorm.DB) Product {
	db.AutoMigrate(&product{})
	mockData(db)
	return &productRepositoryDB{db}
}

func (obj productRepositoryDB) GetProduct() (products []product, err error) {

	// var products []product

	err = obj.db.Order("quantity DESC").Limit(50).Find(&products).Error

	return
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
