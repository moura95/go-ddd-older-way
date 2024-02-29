package main

import (
	"fmt"
	_ "github.com/lib/pq"
	http "go-ddd/cmd/http"
	"go-ddd/internal/infra/cfg"
	"go-ddd/internal/infra/database"
	"go.uber.org/zap"
	"log"
)

func main() {
	// Configs
	loadConfig, _ := cfg.LoadConfig(".")

	// instance Db
	conn, err := database.ConnectPostgres()
	store := conn.DB()
	if err != nil {
		fmt.Println("Failed to Connected Database")
		panic(err)
	}
	log.Print("connection is database establish")

	// Zap Logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	// Run Gin
	http.RunGinServer(loadConfig, store, sugar)
}
