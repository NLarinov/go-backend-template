package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hokamsingh/go-backend-template/internal/config"
	"github.com/hokamsingh/go-backend-template/internal/database"
	"github.com/hokamsingh/go-backend-template/internal/routes"
	"github.com/hokamsingh/go-backend-template/internal/server"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize PostgreSQL
	if err := database.Initialize(database.NewPostgresConfig()); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Get database connection
	db := database.GetDB()
	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get DB connection: %v", err)
	}
	defer sqlDb.Close()

	// Setup router
	router := routes.SetupRouter(db)

	// Use the server abstraction
	srv := server.NewServer(router)

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		log.Println("Shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Server forced to shutdown:", err)
		}
	}()

	fmt.Println("Server started on port", cfg.Server.Port)
	if err := srv.Start(cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
