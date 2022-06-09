package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Gidraff/task-manager-service/config"
	"github.com/Gidraff/task-manager-service/model"
	"github.com/Gidraff/task-manager-service/server"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load configuration parameters
	path, _ := os.Getwd()
	configPath := filepath.Join(path, "/config")
	conf := config.LoadConfig(configPath)

	dbHost := conf.GetString("POSTGRES_HOST")
	username := conf.GetString("POSTGRES_USER")
	password := conf.GetString("POSTGRES_PASSWORD")
	dbName := conf.GetString("POSTGRES_DB")
	//dbPort := conf.GetString("PG_PORT")

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s ", dbHost, username, dbName, password)
	fmt.Println("dns", dsn)

	// Open db connection
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB: Error while establishing connection %+v", err)
	}
	log.Println("DB: connection established")

	db.AutoMigrate(&model.User{}, &model.Project{}, &model.Task{})

	// Run app
	app := server.NewApp(db)
	if err := app.Run(); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
