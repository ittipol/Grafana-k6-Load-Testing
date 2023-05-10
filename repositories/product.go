package repositories

type product struct {
	ID       int
	Name     string `gorm:"size:100"`
	Quantity int
}

type Product interface {
	GetProduct() ([]product, error)
}
