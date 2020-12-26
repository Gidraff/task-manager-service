package main

import (
	"fmt"
	"github.com/Gidraff/task-manager-service/config"
	"github.com/Gidraff/task-manager-service/model"
	"github.com/Gidraff/task-manager-service/server"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Load config file
	path, _ := os.Getwd()
	configPath := filepath.Join(path, "/config")
	conf := config.LoadConfig(configPath)
	var dsn string

	if conf.GetString("ENV") != "dev" {
		dsn = fmt.Sprintf(
			conf.GetString("DB_DSN"),
			conf.GetString("DB_USER"),
			conf.GetString("DB_NAME"),
			conf.GetString("DB_PASSWORD"),
			conf.GetString("DB_PORT"),
		)
	} else {
		dsn = fmt.Sprintf(
			conf.GetString("dsn"),
			conf.GetString("database.user"),
			conf.GetString("database.name"),
			conf.GetString("database.password"),
			conf.GetString("database.port"),
		)
	}

	// Open db connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB: Error while establishing connection %+v", err)
	}
	log.Println("DB: connection established")

	db.AutoMigrate(&model.User{}, &model.Project{}, &model.Task{})

	// Run app
	app := server.NewApp(db)
	if err := app.Run(conf.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
