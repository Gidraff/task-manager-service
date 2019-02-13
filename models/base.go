package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	// Use postgres
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

// Db connections handler
var db *gorm.DB

func init() {
	e := godotenv.Load()
	if e != nil {
		log.Fatal("Error loading .env file")
	}

	dbname := os.Getenv("PG_DB")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")

	connectionString := fmt.Sprintf(
		"host=localhost port=5432 user=%s dbname=%s password=%s sslmode=disable",
		user, dbname, password)

	dbConnection, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic(err.Error())
	}
	db = dbConnection
	log.Printf("Connection Established")
	GetDB().Debug().AutoMigrate(&Account{}, &Token{})
}

// GetDB return db connection handle
func GetDB() *gorm.DB {
	return db
}
