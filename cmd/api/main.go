package main

import (
	"database/sql"
	"fmt"
	"github.com/Gidraff/task-manager-service/config"
	"github.com/Gidraff/task-manager-service/pkg/utils/logger"
	"github.com/Gidraff/task-manager-service/server"
	_ "github.com/lib/pq"
	"log"
	"os"
	"path/filepath"
)

func main() {

	// Load config file
	path, _ := os.Getwd()
	configPath := filepath.Join(path, "/config")
	conf := config.LoadConfig(configPath)

	// Create connection string
	connStr := fmt.Sprintf(
		conf.GetString("dsn"),
		conf.GetString("database.user"),
		conf.GetString("database.password"),
		conf.GetString("database.name"),
	)

	// Open db connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("DB: Error while establishing connection %+v", err)
	}

	defer db.Close()

	// Run app
	logger := logger.NewLogger("info")
	app := server.NewApp(db, logger)
	if err := app.Run(conf.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
