package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Gidraff/task-manager-service/handlers"
	"github.com/Gidraff/task-manager-service/middlewares"
	"github.com/Gidraff/task-manager-service/models"
	"github.com/gorilla/mux"
)

var wait time.Duration

func main() {
	err := models.GetDB().DB().Ping()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	defer models.GetDB().Close()

	// http server setup
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "15s")
	flag.Parse()

	router := mux.NewRouter()
	router.Use(middlewares.JwtAuthentication) // attach middleware

	router.HandleFunc("/api/user/new", handlers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", handlers.Authenticate).Methods("POST")

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// accept graceful shutdowns when quit via SIGINT(Ctrl+C)
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Does not block if no connections, but will otherwise wait
	// until the timeout deadline
	server.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
