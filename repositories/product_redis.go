package repositories

type productRepositoryRedis struct {
	// Product
	db interface{}
}

func NewProductRepositoryRedis(db interface{}) Product {
	return &productRepositoryRedis{db}
}

func (obj productRepositoryRedis) GetProduct() ([]product, error) {
	return nil, nil
}
