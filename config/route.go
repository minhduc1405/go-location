package config

import (
	"context"
	"github.com/go-chi/chi"
)

func Route(r *chi.Mux, context context.Context, config Root) error {
	applicationContext, err := NewApplicationContext(context, config)
	if err != nil {
		return err
	}

	locationController := applicationContext.LocationController
	locationPath := "/locations"
	// r.Get(locationPath, locationController.All)
	r.Get(locationPath+"", locationController.Search)
	r.Post(locationPath+"/search", locationController.Search)
	r.Get(locationPath+"/{id}", locationController.Load)
	r.Post(locationPath, locationController.Insert)
	r.Put(locationPath+"/:locationId", locationController.Update)

	return nil
}
