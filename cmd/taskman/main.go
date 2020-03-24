package main

import (
	"database/sql"
	"github.com/getsentry/sentry-go"
	"log"
	"time"

	"github.com/Gidraff/task-manager-service/cmd/taskman/config"
	_userHttpDeliver "github.com/Gidraff/task-manager-service/cmd/taskman/delivery/http"

	_userRepo "github.com/Gidraff/task-manager-service/cmd/taskman/repository"
	_userService "github.com/Gidraff/task-manager-service/cmd/taskman/service"
	"github.com/labstack/echo"

	// _ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"

	"fmt"
	"os"
	"path/filepath"
	// "github.com/Gidraff/go-interfaces/cmd/taskman/config"
)

func main() {

	err := sentry.Init(sentry.ClientOptions{
		Dsn:   "http://220e0359f4dc4ff090bbf5ea7f4cb644@sentry.io/4438912",
		Debug: true,
	})

	if err != nil {
		log.Fatalf("Sentry.Init: %s", err)
	}

	defer sentry.Flush(2 * time.Second)
	path, _ := os.Getwd()
	configPath := filepath.Join(path, "/config")
	cfg := config.LoadConfig(configPath)

	connStr := fmt.Sprintf(
		cfg.GetString("dsn"),
		cfg.GetString("database.dbuser"),
		cfg.GetString("database.dbpassword"),
		cfg.GetString("database.dbname"),
	)

	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		sentry.CaptureException(err)
		log.Printf("Failed to make connection: %s", err)
	}

	defer dbConn.Close()

	// Initialize router
	e := echo.New()
	userRepo := _userRepo.NewUserRepo(dbConn)
	uService := _userService.NewService(userRepo)
	_userHttpDeliver.NewUserHandler(e, uService)
	log.Fatal(e.Start(cfg.GetString("port")))

}
