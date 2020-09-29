package repository

import "go-service/internal/model"

type MerchantRepository interface {
	Create(model *model.Merchant) (*model.Merchant, error)
	Update(model *model.Merchant) error
	Delete(id string) error
	All() ([]model.Merchant, error)
	Load(id string) (*model.Merchant, error)
}
