package domain

type Core struct {
	ID          uint
	Name        string
	Description string
	Qty         int
	Price       int
}

type Repository interface {
	ShowAll(category, name string, page int) ([]Core, error)
	Insert(newProduct Core) (Core, error)
	Delete(ID uint) error
}

type Services interface {
	GetAll(category, name string, page int) ([]Core, error)
	AddProduct(newProduct Core) (Core, error)
	Destroy(ID uint) error
}
