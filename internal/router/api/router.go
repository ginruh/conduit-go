package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/iyorozuya/real-world-app/internal/controllers/api"
	serviceApi "github.com/iyorozuya/real-world-app/internal/services/api"
	"github.com/iyorozuya/real-world-app/internal/utils"
)

func Bootstrap(r chi.Router, db utils.DB) {
	// User login and registration
	r.Group(func(r chi.Router) {
		authController := api.AuthController{
			AuthService: serviceApi.AuthServiceImpl{
				DB: &db,
			},
		}
		r.Post("/users/login", authController.Login)
		r.Post("/users", authController.Register)
	})
}
