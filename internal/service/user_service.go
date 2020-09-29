package service

import (
	"go-service/internal/model"
	"go-service/internal/repository"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserService {
	return &UserService{repository}
}

func (s *UserService) Create(model *model.User) (*model.User, error) {
	return s.repository.Create(model)
}

func (s *UserService) Update(model *model.User) error {
	return s.repository.Update(model)
}

func (s *UserService) Delete(id string) error {
	return s.repository.Delete(id)
}

func (s *UserService) All() ([]model.User, error) {
	return s.repository.All()
}

func (s *UserService) Load(id string) (*model.User, error) {
	return s.repository.Load(id)
}
