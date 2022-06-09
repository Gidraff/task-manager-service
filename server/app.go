package server

import (
	"context"

	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Gidraff/task-manager-service/pkg/utils/helpers"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/Gidraff/task-manager-service/auth"
	authHTTP "github.com/Gidraff/task-manager-service/auth/delivery/http"
	authPostgres "github.com/Gidraff/task-manager-service/auth/repository/postgres"
	authUsecase "github.com/Gidraff/task-manager-service/auth/usecase"
	"github.com/Gidraff/task-manager-service/project"
	projectHTTP "github.com/Gidraff/task-manager-service/project/delivery/http"
	projectPostgres "github.com/Gidraff/task-manager-service/project/repository/postgres"
	projectUsecase "github.com/Gidraff/task-manager-service/project/usecase"
	"github.com/Gidraff/task-manager-service/task"
	taskHTTP "github.com/Gidraff/task-manager-service/task/delivery/http"
	taskPostgres "github.com/Gidraff/task-manager-service/task/repository/postgres"
	taskUsecase "github.com/Gidraff/task-manager-service/task/usecase"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// App encapsulate application' server
type App struct {
	httpServer *http.Server
	authUC     auth.UseCase
	taskUC     task.ITaskUsecase
	projectUC  project.IProjectUsecase
}

// NewApp return a instance of app
func NewApp(db *gorm.DB) *App {
	userAuthRepo := authPostgres.NewUserRepo(db)
	projectRepo := projectPostgres.NewProjectRepo(db)
	taskRepo := taskPostgres.NewTaskRepo(db)
	return &App{
		authUC:    authUsecase.NewUseCase(userAuthRepo),
		taskUC:    taskUsecase.NewTaskUsecase(taskRepo),
		projectUC: projectUsecase.NewProjectUsecase(projectRepo),
	}
}

// Run bootstraps the app's server
func (a *App) Run() error {
	router := mux.NewRouter().StrictSlash(false)
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(router)

	router.HandleFunc("/api/v1/status", func(w http.ResponseWriter, r *http.Request) {
		helpers.Response(http.StatusOK, "Up!", w)
	}).Methods("GET")
	//authR := router.PathPrefix("/api/v1/auth").Subrouter()
	authMiddleware := mux.NewRouter()
	router.PathPrefix("/api/v1/").Handler(negroni.New(
		negroni.HandlerFunc(helpers.JwtMiddleware),
		negroni.Wrap(authMiddleware),
	))

	// Register Endpoints
	authHTTP.RegisterHandler(authMiddleware, a.authUC)
	taskHTTP.RegisterHandler(authMiddleware, a.taskUC)
	projectHTTP.RegisterHandler(authMiddleware, a.projectUC)

	a.httpServer = &http.Server{
		Addr:           ":8089",
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
