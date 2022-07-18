package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/iyorozuya/real-world-app/docs"
	"github.com/iyorozuya/real-world-app/internal/db"
	"github.com/iyorozuya/real-world-app/internal/router/api"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
)

// @title    Conduit
// @version  1.0
// description Medium clone according to real-world-app specs

// @host      localhost:5000
// @BasePath  /api
func main() {
	// Load from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to db
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbInstance := db.New(dbUsername, dbPassword, dbHost, dbPort, dbName)
	db, err := dbInstance.Connect()
	if err != nil {
		log.Fatalln("Unable to connect to db", err)
		return
	}
	log.Println("Connected to db successfully")

	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.StripSlashes)

	r.Route("/api", func(r chi.Router) {
		// Swagger setup
		r.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL(fmt.Sprintf("http://localhost:%v/api/swagger/doc.json", os.Getenv("API_PORT"))),
		))
		// bootstrapping api
		api.Bootstrap(r, db)
	})

	log.Printf("Server is running at port %v", os.Getenv("API_PORT"))
	if err := http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("API_PORT")), r); err != nil {
		log.Fatalln(err)
	}
}
