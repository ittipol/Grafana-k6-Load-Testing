package repositories

import (
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

func (obj productRepositoryDB) GetCachedProduct() (products []product, err error) {
	return nil, nil
}
