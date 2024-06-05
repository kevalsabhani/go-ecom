package types

type UserStore interface {
	GetUserByEmail(string) (*User, error)
	GetUserByID(string) (*User, error)
	CreateUser(*User) error
}
