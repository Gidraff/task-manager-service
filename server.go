package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	// Use postgres
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

// Db connection pool
var Db *gorm.DB
var err error
var wait time.Duration

// Load .env file
func init() {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbname := os.Getenv("POSTGRES_DB")
	user := os.Getenv("POSTGRES_USER")
	host := os.Getenv("HOST")
	password := os.Getenv("POSTGRES_PASSWORD")

	connectionString := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s sslmode=disable",
		host, user, dbname, password)

	Db, err = gorm.Open("postgres", connectionString)
	if err != nil {
		panic(err.Error())
	}
	log.Printf("Successfully connected to db")
}

func main() {
	err = Db.DB().Ping()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	defer Db.Close()

	// http server setup
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "15s")
	flag.Parse()

	r := mux.NewRouter()

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
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
