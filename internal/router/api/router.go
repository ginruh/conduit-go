package api

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/iyorozuya/real-world-app/internal/controllers/api"
	"github.com/iyorozuya/real-world-app/internal/middlewares"
	"github.com/iyorozuya/real-world-app/internal/seeders"
	"github.com/iyorozuya/real-world-app/internal/services/api/auth"
	"github.com/iyorozuya/real-world-app/internal/services/api/user"
	"github.com/iyorozuya/real-world-app/internal/sqlc"
	"os"
)

func Bootstrap(r chi.Router, database *sql.DB) {
	q := sqlc.New(database)
	validate := validator.New()

	authController := api.NewAuthController(auth.NewAuthService(q), validate)
	userController := api.NewUserController(user.NewUserService(q), validate)

	// User login and registration
	r.Group(func(r chi.Router) {
		r.Post("/users/login", authController.Login)
		r.Post("/users", authController.Register)
	})

	// User authenticated actions
	r.Group(func(r chi.Router) {
		r.Use(middlewares.JWT)
		r.Get("/user", userController.GetCurrentUser)
		r.Put("/user", userController.UpdateUser)
	})

	// Optional authenticated routes
	r.Group(func(r chi.Router) {
		r.Use(middlewares.OptionalJWT)
		r.Get("/profile/{username}", userController.GetProfile)
	})

	// Seed db if RUN_SEEDERS is yes
	if os.Getenv("RUN_SEEDERS") == "yes" {
		seed := seeders.New(q)
		seed.SeedAll()
	}
}
