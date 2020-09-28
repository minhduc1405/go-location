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

	// user
	userController := applicationContext.UserController
	userParent := "/users"
	r.Post(userParent, userController.CreateUser)
	r.Patch(userParent+"/{id}", userController.UpdateUser)
	r.Delete(userParent+"/{id}", userController.DeleteUser)
	r.Get(userParent, userController.GetAllUsers)
	r.Get(userParent+"/{id}", userController.LoadUser)

	// merchant
	merchantController := applicationContext.MerchantController
	merchantParent := "/merchants"
	r.Post(merchantParent, merchantController.CreateMerchant)
	r.Patch(merchantParent+"/{id}", merchantController.UpdateMerchant)
	r.Delete(merchantParent+"/{id}", merchantController.DeleteMerchant)
	r.Get(merchantParent, merchantController.GetAllMerchants)
	r.Get(merchantParent+"/{id}", merchantController.LoadMerchant)
	return nil
}
