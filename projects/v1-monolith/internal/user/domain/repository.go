package domain

type UserRepository interface {
	FindById(ID string) (*User, error)
	FindMany() ([]User, error)
	Create(user *User) error
}
