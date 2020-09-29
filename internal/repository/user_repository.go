package repository

import "go-service/internal/model"

type UserRepository interface {
	Create(model *model.User) (*model.User, error)
	Update(model *model.User) error
	Delete(id string) error
	All() ([]model.User, error)
	Load(id string) (*model.User, error)
}
