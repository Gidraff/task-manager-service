package main

import (
	"github.com/Gidraff/task-manager-service/config"
	"github.com/Gidraff/task-manager-service/server"
	_ "github.com/lib/pq"
	"log"
	"os"
	"path/filepath"
)

func main() {
	path, _ := os.Getwd()
	configPath := filepath.Join(path, "/config")
	cfg := config.LoadConfig(configPath)
	app := server.NewApp(cfg)
	if err := app.Run(cfg.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
