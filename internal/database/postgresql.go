// internal/database/postgresql.go
package database

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/hokamsingh/go-backend-template/internal/config"
	"github.com/hokamsingh/go-backend-template/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbOnce sync.Once
	db     *gorm.DB
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresConfig() PostgresConfig {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	return PostgresConfig{
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		DBName:   cfg.Database.Name,
		SSLMode:  cfg.Database.SSLMode,
	}
}

func NewPostgresConnection(config PostgresConfig) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	maxRetries := 5
	retryDelay := 5 * time.Second

	for i := 0; i < maxRetries; i++ {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode)

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}

		log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
		if i < maxRetries-1 {
			log.Printf("Retrying in %v...", retryDelay)
			time.Sleep(retryDelay)
		}
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database after %d attempts: %w", maxRetries, err)
	}

	// Run migrations
	if err := RunMigrations(db); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	// Seed database with initial data
	if err := SeedDatabase(db); err != nil {
		log.Printf("Warning: failed to seed database: %v", err)
	}

	return db, nil
}

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Event{},
		&models.Speaker{},
		&models.Tag{},
	)
}

// Initialize initializes database connection
func Initialize(config PostgresConfig) error {
	var err error
	dbOnce.Do(func() {
		db, err = NewPostgresConnection(config)
	})
	return err
}

// GetDB returns the PostgreSQL database instance
func GetDB() *gorm.DB {
	return db
}
