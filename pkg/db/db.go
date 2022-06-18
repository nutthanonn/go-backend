package db

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
	"github.com/nutthanonn/golang-webserver/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var projectDirName = "golang-webserver"

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func Init() *gorm.DB {
	loadEnv()
	dbURL := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("PORT"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Year{})

	return db

}
