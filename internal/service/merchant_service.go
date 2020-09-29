package service

import (
	"go-service/internal/model"
	"go-service/internal/repository"
)

type MerchantService struct {
	repository repository.MerchantRepository
}

func NewMerchantService(repository repository.MerchantRepository) *MerchantService {
	return &MerchantService{repository}
}

func (s *MerchantService) Create(model *model.Merchant) (*model.Merchant, error) {
	return s.repository.Create(model)
}

func (s *MerchantService) Update(model *model.Merchant) error {
	return s.repository.Update(model)
}

func (s *MerchantService) Delete(id string) error {
	return s.repository.Delete(id)
}

func (s *MerchantService) All() ([]model.Merchant, error) {
	return s.repository.All()
}

func (s *MerchantService) Load(id string) (*model.Merchant, error) {
	return s.repository.Load(id)
}
