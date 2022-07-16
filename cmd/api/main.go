package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/iyorozuya/real-world-app/docs"
	"github.com/iyorozuya/real-world-app/internal/router/api"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

// @title    Conduit
// @version  1.0
// description Medium clone according to real-world-app specs

// @host      localhost:5000
// @BasePath  /api
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		// Swagger setup
		r.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL("http://localhost:5000/api/swagger/doc.json"), //The url pointing to API definition
		))

		// bootstrapping api
		api.Bootstrap(r)
	})

	http.ListenAndServe(":5000", r)
}
