package types

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(user *User) error
}
