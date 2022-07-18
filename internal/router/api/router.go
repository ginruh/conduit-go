package api

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/iyorozuya/real-world-app/internal/controllers/api"
	"github.com/iyorozuya/real-world-app/internal/seeders"
	serviceApi "github.com/iyorozuya/real-world-app/internal/services/api/auth"
	"github.com/iyorozuya/real-world-app/internal/sqlc"
	"os"
)

func Bootstrap(r chi.Router, database *sql.DB) {
	q := sqlc.New(database)
	validate := validator.New()

	// User login and registration
	authController := api.NewAuthController(serviceApi.NewAuthService(q), validate)
	r.Post("/users/login", authController.Login)
	r.Post("/users", authController.Register)

	// Seed db if RUN_SEEDERS is yes
	if os.Getenv("RUN_SEEDERS") == "yes" {
		seed := seeders.New(q)
		seed.SeedAll()
	}
}
