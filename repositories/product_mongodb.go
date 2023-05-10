package repositories

type productRepositoryMongoDB struct {
	Product
	// db interface{}
}

func NewProductRepositoryMongoDB() Product {
	return &productRepositoryMongoDB{}
}

func (obj productRepositoryMongoDB) GetProduct() ([]product, error) {
	return nil, nil
}
