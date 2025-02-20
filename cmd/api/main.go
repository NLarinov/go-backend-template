package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/hokamsingh/go-backend-template/internal/config"
	"github.com/hokamsingh/go-backend-template/internal/database"
	"github.com/hokamsingh/go-backend-template/internal/handlers"
	"github.com/hokamsingh/go-backend-template/internal/repository"
	"github.com/hokamsingh/go-backend-template/internal/service"
	"github.com/redis/go-redis/v9"
)

func main() {
	cfg, _ := config.LoadConfig()

	db, err := database.NewPostgresConnection(database.NewPostgresConfig())
	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %v", err)
	}
	sqlDb, _ := db.DB()
	defer sqlDb.Close()

	redisClient := initRedis(cfg)
	defer redisClient.Close()

	userService := service.NewUserService(repository.NewUserRepository(db))
	router := gin.Default()
	registerRoutes(router, userService)

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		fmt.Println("Shutting down server...")
		redisClient.Close()
		os.Exit(0)
	}()

	port := fmt.Sprintf(":%s", cfg.Server.Port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func initRedis(cfg *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	// Check if Redis is reachable
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	return client
}

func registerRoutes(router *gin.Engine, userService *service.UserService) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/create", handlers.CreateUser(userService))
		userRoutes.GET("/:id", handlers.GetUser(userService))
	}
}
