package utils

import (
	"database/sql"
	"github.com/iyorozuya/real-world-app/internal/db"
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
)

func ConnectTestDB() *sql.DB {
	projectDirName := "real-world-app"
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))
	if err := godotenv.Load(string(rootPath) + `/.env`); err != nil {
		log.Fatalln("Unable to load .env")
	}
	dbInstance := db.New(
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_TEST_NAME"),
	)
	db, err := dbInstance.Connect()
	if err != nil {
		log.Fatalln("Unable to connect to test database", err)
	}
	return db
}
