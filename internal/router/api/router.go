package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/iyorozuya/real-world-app/internal/controllers/api"
	serviceApi "github.com/iyorozuya/real-world-app/internal/services/api"
)

func Bootstrap(r chi.Router) {
	// User login and registration
	r.Group(func(r chi.Router) {
		authController := api.AuthController{
			AuthService: serviceApi.AuthServiceImpl{},
		}
		r.Post("/users/login", authController.Login)
		r.Post("/users", authController.Register)
	})
}
