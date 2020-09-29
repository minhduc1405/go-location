package location

import (
	m "github.com/common-go/mongo"
	. "github.com/common-go/search"
	. "github.com/common-go/service"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
)

type MongoLocationService struct {
	GenericService
	SearchService
	LocationMapper m.Mapper
}

func NewMongoLocationService(db *mongo.Database, searchBuilder m.SearchResultBuilder, mapper m.Mapper) *MongoLocationService {
	var model Location
	typeOfModel := reflect.TypeOf(model)
	genericService, searchService := m.NewGenericSearchService(db, typeOfModel, "location", searchBuilder, false, "Version", mapper)
	return &MongoLocationService{genericService, searchService, mapper}
}
