package app

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go-service/internal/handler"
	"go-service/internal/repository"
	"go-service/internal/service"
)

type ApplicationContext struct {
	UserHandler     *handler.UserHandler
	MerchantHandler *handler.MerchantHandler
}

func NewApplicationContext(context context.Context, config Root) (*ApplicationContext, error) {
	sqlUri := config.Sql.Uri
	sqlDb, err := gorm.Open(mysql.Open(sqlUri), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	//sqlDb.AutoMigrate(merchant.Merchant{})
	userRepository := repository.NewUserRepository(sqlDb, "users")
	userService := service.NewUserService(userRepository)
	userController := handler.NewUserHandler(userService)

	merchantRepository := repository.NewMerchantRepository(sqlDb, "merchants")
	merchantService := service.NewMerchantService(merchantRepository)
	merchantController := handler.NewMerchantHandler(merchantService)

	app := &ApplicationContext{
		UserHandler:     userController,
		MerchantHandler: merchantController,
	}
	return app, nil
}
