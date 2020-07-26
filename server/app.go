package server

import (
	"context"
	log "github.com/sirupsen/logrus"

	//"database/sql"
	"fmt"
	"github.com/Gidraff/task-manager-service/auth"
	authhttp "github.com/Gidraff/task-manager-service/auth/delivery/http"
	authpostgres "github.com/Gidraff/task-manager-service/auth/repository/postgres"
	authusecase "github.com/Gidraff/task-manager-service/auth/usecase"
	lg "github.com/Gidraff/task-manager-service/pkg/utils/logger"
	"github.com/gorilla/mux"
	_ "github.com/hpcloud/tail/util"
	"github.com/jinzhu/gorm"
	"github.com/urfave/negroni"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// App encapsulate application' server
type App struct {
	httpServer *http.Server
	authUC     auth.UseCase
	logger     *lg.Logger
}

// NewApp return a instance of app
func NewApp(db *gorm.DB, logger *lg.Logger) *App {
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
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}
