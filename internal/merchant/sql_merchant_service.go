package merchant

import (
	"errors"
	"gorm.io/gorm"
)

type SqlMerchantService struct {
	db        *gorm.DB
	tableName string
}

func NewMerchantService(db *gorm.DB, tableName string) (*SqlMerchantService, error) {
	if len(tableName) < 1 {
		return nil, errors.New("table name cannot be empty")
	}
	return &SqlMerchantService{db, tableName}, nil
}

func (s *SqlMerchantService) Create(model *Merchant) (*Merchant, error) {
	err := s.db.Create(model).Table(s.tableName).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (s *SqlMerchantService) Update(model *Merchant) error {
	err := s.db.Model(model).Updates(model).Table(s.tableName).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlMerchantService) Delete(id string) error {
	err := s.db.Unscoped().Delete(&Merchant{}, id).Table(s.tableName).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlMerchantService) All() ([]Merchant, error) {
	var rs []Merchant
	err := s.db.Unscoped().Find(&rs).Table(s.tableName).Error
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func (s *SqlMerchantService) Load(id string) (*Merchant, error) {
	var rs Merchant
	err := s.db.Unscoped().First(&rs, id).Table(s.tableName).Error
	if err != nil {
		return nil, err
	}
	return &rs, nil
}
