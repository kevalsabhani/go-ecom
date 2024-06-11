package types

type UserStore interface {
	GetUserByEmail(string) (*User, error)
	GetUserByID(string) (*User, error)
	CreateUser(*User) error
}

type ProductStore interface {
	GetProducts() ([]*Product, error)
	CreateProduct(*Product) error
}
