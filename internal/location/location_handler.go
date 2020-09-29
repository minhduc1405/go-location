package location

import (
	"github.com/common-go/service"
	"github.com/common-go/validator"
	"github.com/common-go/web"
	"reflect"
)

type LocationHandler struct {
	*server.GenericController
	*server.SearchController
}

func NewLocationHandler(locationService LocationService, validator validator.Validator, logService server.ActivityLogService) *LocationHandler {
	modelType := reflect.TypeOf(Location{})
	searchModelType := reflect.TypeOf(LocationSM{})
	idGenerator := service.NewIdGenerator(false, false, false)
	modelBuilder := service.NewModelBuilder(idGenerator, modelType, "CreatedBy", "CreatedAt", "UpdatedBy", "UpdatedAt")
	genericController, searchController := server.NewGenericSearchController(locationService, modelType, modelBuilder, locationService, searchModelType, validator, logService, false, "", "")
	return &LocationHandler{genericController, searchController}
}
