package app

import (
	"context"
	m "github.com/common-go/mongo"
	"github.com/common-go/validator"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"reflect"

	"go-location/internal/location"
	"go-location/internal/merchant"
	"go-location/internal/user"
)

type ApplicationContext struct {
	LocationHandler *location.LocationHandler
	UserHandler     *user.UserHandler
	MerchantHandler *merchant.MerchantHandler
}

func NewApplicationContext(context context.Context, config Root) (*ApplicationContext, error) {
	client, err := mongo.Connect(context, options.Client().ApplyURI(config.Mongo.Uri))
	if err != nil {
		return nil, err
	}

	db := client.Database(config.Mongo.Database)

	validator := validator.NewDefaultValidator()
	queryBuilder := &m.DefaultQueryBuilder{}
	sortBuilder := &m.DefaultSortBuilder{}
	modelType := reflect.TypeOf(location.Location{})
	locationMapper := m.NewMapper(modelType, "Latitude", "Longitude", "Location")
	mongoSearchResultBuilder := &m.DefaultSearchResultBuilder{
		Database:     db,
		QueryBuilder: queryBuilder,
		SortBuilder:  sortBuilder,
		Mapper:       locationMapper,
	}

	locationService := location.NewMongoLocationService(db, mongoSearchResultBuilder, locationMapper)
	locationController := location.NewLocationHandler(locationService, validator, nil)

	//merchant
	sqlUri := config.Sql.Uri
	sqlDb, er0 := gorm.Open(mysql.Open(sqlUri), &gorm.Config{})
	if er0 != nil {
		return nil, er0
	}
	//sqlDb.AutoMigrate(merchant.Merchant{})
	userService, er1 := user.NewUserService(sqlDb, "users")
	if er1 != nil {
		return nil, er1
	}
	userController := user.NewUserHandler(userService)

	merchantService, er2 := merchant.NewMerchantService(sqlDb, "merchants")
	if er2 != nil {
		return nil, er2
	}
	merchantController := merchant.NewMerchantHandler(merchantService)

	app := &ApplicationContext{
		LocationHandler: locationController,
		UserHandler:     userController,
		MerchantHandler: merchantController,
	}
	return app, nil
}
