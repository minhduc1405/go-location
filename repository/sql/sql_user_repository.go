package sql

import (
	"go-service/model"
	"gorm.io/gorm"
)

type SqlUserRepository struct {
	db        *gorm.DB
	tableName string
}

func NewUserRepository(db *gorm.DB, tableName string) *SqlUserRepository {
	return &SqlUserRepository{db, tableName}
}

func (s *SqlUserRepository) Create(model *model.User) (*model.User, error) {

	err := s.db.Create(model).Table(s.tableName).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *SqlUserRepository) Update(model *model.User) error {
	err := s.db.Model(model).Updates(model).Table(s.tableName).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlUserRepository) Delete(id string) error {
	err := s.db.Unscoped().Delete(&model.User{}, id).Table(s.tableName).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlUserRepository) All() ([]model.User, error) {
	var rs []model.User
	err := s.db.Unscoped().Find(&rs).Table(s.tableName).Error
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func (s *SqlUserRepository) Load(id string) (*model.User, error) {
	var rs model.User
	err := s.db.Unscoped().First(&rs, id).Table(s.tableName).Error
	if err != nil {
		return nil, err
	}
	return &rs, nil
}
