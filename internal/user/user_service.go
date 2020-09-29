package user

type UserService interface {
	Create(model *User) (*User, error)
	Update(model *User) error
	Delete(id string) error
	All() ([]User, error)
	Load(id string) (*User, error)
}
