package server

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Gidraff/task-manager-service/auth"
	authhttp "github.com/Gidraff/task-manager-service/auth/delivery/http"
	authpostgres "github.com/Gidraff/task-manager-service/auth/repository/postgres"
	authusecase "github.com/Gidraff/task-manager-service/auth/usecase"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// App encapsulate application' server
type App struct {
	httpServer *http.Server
	authUC     auth.BasicAuthUseCase
}

// NewApp return a instance of app
func NewApp(cfg *viper.Viper) *App {
	db := initDB(cfg)
	defer db.Close()

	userAuthRepo := authpostgres.NewUserRepo(db)
	return &App{
		authUC: authusecase.NewUseCase(userAuthRepo),
	}

}

// Run bootstraps the app's server
func (a *App) Run(port string) error {
	router := mux.NewRouter()
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(router)

	// Register Endpoints
	authhttp.RegisterHttpEndpoints(router, a.authUC)

	a.httpServer = &http.Server{
		Addr:           port,
		Handler:        n,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		fmt.Println("Server starting...")
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and server: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

// initDB return a db connection
func initDB(vc *viper.Viper) *sql.DB {
	connStr := fmt.Sprintf(
		vc.GetString("dsn"),
		vc.GetString("database.dbuser"),
		vc.GetString("database.dbpassword"),
		vc.GetString("database.dbname"),
	)

	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("An error occured while trying to connect %+v", err)
	}
	return dbConn
}
