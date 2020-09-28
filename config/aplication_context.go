package config

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go-service/controller"
	"go-service/repository/sql"
	"go-service/service"
)

type ApplicationContext struct {
	UserController     *controller.UserController
	MerchantController *controller.MerchantController
}

func NewApplicationContext(context context.Context, config Root) (*ApplicationContext, error) {
	sqlUri := config.Sql.Uri
	sqlDb, err := gorm.Open(mysql.Open(sqlUri), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	//sqlDb.AutoMigrate(merchant.Merchant{})
	userRepository := sql.NewUserRepository(sqlDb, "users")
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	merchantRepository := sql.NewMerchantRepository(sqlDb, "merchants")
	merchantService := service.NewMerchantService(merchantRepository)
	merchantController := controller.NewMerchantController(merchantService)

	app := &ApplicationContext{
		UserController:     userController,
		MerchantController: merchantController,
	}
	return app, nil
}
