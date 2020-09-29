package user

import (
	"errors"
	"gorm.io/gorm"
)

type SqlUserService struct {
	db        *gorm.DB
	tableName string
}

func NewUserService(db *gorm.DB, tableName string) (*SqlUserService, error) {
	if len(tableName) < 1 {
		return nil, errors.New("table name cannot be empty")
	}
	return &SqlUserService{db, tableName}, nil
}

func (s *SqlUserService) Create(model *User) (*User, error) {

	err := s.db.Create(model).Table(s.tableName).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *SqlUserService) Update(model *User) error {
	err := s.db.Model(model).Updates(model).Table(s.tableName).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlUserService) Delete(id string) error {
	err := s.db.Unscoped().Delete(&User{}, id).Table(s.tableName).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlUserService) All() ([]User, error) {
	var rs []User
	err := s.db.Unscoped().Find(&rs).Table(s.tableName).Error
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func (s *SqlUserService) Load(id string) (*User, error) {
	var rs User
	err := s.db.Unscoped().First(&rs, id).Table(s.tableName).Error
	if err != nil {
		return nil, err
	}
	return &rs, nil
}
