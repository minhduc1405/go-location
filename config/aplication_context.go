package config

import (
	"context"
	m "github.com/common-go/mongo"
	"github.com/common-go/validator"
	"go-location/location"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
)

type ApplicationContext struct {
	LocationController *location.LocationController
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
	locationController := location.NewLocationController(locationService, validator, nil)
	app := &ApplicationContext{
		LocationController: locationController,
	}
	return app, nil
}
