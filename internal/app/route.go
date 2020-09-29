package app

import (
	"context"
	"github.com/go-chi/chi"
)

func Route(r *chi.Mux, context context.Context, config Root) error {
	applicationContext, err := NewApplicationContext(context, config)
	if err != nil {
		return err
	}

	// user
	userHandler := applicationContext.UserHandler
	userParent := "/users"
	r.Post(userParent, userHandler.CreateUser)
	r.Patch(userParent+"/{id}", userHandler.UpdateUser)
	r.Delete(userParent+"/{id}", userHandler.DeleteUser)
	r.Get(userParent, userHandler.GetAllUsers)
	r.Get(userParent+"/{id}", userHandler.LoadUser)

	// merchant
	merchantHandler := applicationContext.MerchantHandler
	merchantParent := "/merchants"
	r.Post(merchantParent, merchantHandler.CreateMerchant)
	r.Patch(merchantParent+"/{id}", merchantHandler.UpdateMerchant)
	r.Delete(merchantParent+"/{id}", merchantHandler.DeleteMerchant)
	r.Get(merchantParent, merchantHandler.GetAllMerchants)
	r.Get(merchantParent+"/{id}", merchantHandler.LoadMerchant)
	return nil
}
