package repositories

type productRepositoryMongoDB struct {
	// Product
	db interface{}
}

func NewProductRepositoryMongoDB(db interface{}) Product {
	return &productRepositoryMongoDB{db}
}

func (obj productRepositoryMongoDB) GetProduct() ([]product, error) {
	return nil, nil
}
