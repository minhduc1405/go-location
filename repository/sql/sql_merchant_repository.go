package sql

import (
	"go-service/model"
	"gorm.io/gorm"
)

type SqlMerchantRepository struct {
	db        *gorm.DB
	tableName string
}

func NewMerchantRepository(db *gorm.DB, tableName string) *SqlMerchantRepository {
	return &SqlMerchantRepository{db, tableName}
}

func (s *SqlMerchantRepository) Create(model *model.Merchant) (*model.Merchant, error) {

	err := s.db.Create(model).Table(s.tableName).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *SqlMerchantRepository) Update(model *model.Merchant) error {
	err := s.db.Model(model).Updates(model).Table(s.tableName).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlMerchantRepository) Delete(id string) error {
	err := s.db.Unscoped().Delete(&model.Merchant{}, id).Table(s.tableName).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlMerchantRepository) All() ([]model.Merchant, error) {
	var rs []model.Merchant
	err := s.db.Unscoped().Find(&rs).Table(s.tableName).Error
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func (s *SqlMerchantRepository) Load(id string) (*model.Merchant, error) {
	var rs model.Merchant
	err := s.db.Unscoped().First(&rs, id).Table(s.tableName).Error
	if err != nil {
		return nil, err
	}
	return &rs, nil
}
