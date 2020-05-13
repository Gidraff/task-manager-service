package main

import (
	"database/sql"
	"github.com/Gidraff/task-manager-service/auth/repository/postgres"
	"github.com/getsentry/sentry-go"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"time"

	_authHttpDeliver "github.com/Gidraff/task-manager-service/auth/delivery/http"
	"github.com/Gidraff/task-manager-service/config"

	_authUseCase "github.com/Gidraff/task-manager-service/auth/usecase"

	_ "github.com/lib/pq"

	"fmt"
	"os"
	"path/filepath"
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
	router := mux.NewRouter()
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(router)

	authRepo := postgres.NewAuthRepo(dbConn)
	authUseCase := _authUseCase.NewAuthUseCase(authRepo)
	_authHttpDeliver.NewAuthHandler(router, authUseCase)

	fmt.Println("Server starting...")
	http.ListenAndServe(cfg.GetString("port"), n)
}
