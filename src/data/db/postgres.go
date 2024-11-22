package db

import (
	"fmt"
	"log"
	"time"

	"github.com/mfaxmodem/web-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// db.go
var dbClient *gorm.DB

func InitDb(cfg *config.Config) error {
	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password,
		cfg.Postgres.DbName, cfg.Postgres.SSlMode)

	var err error
	dbClient, err = gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDb, err := dbClient.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}

	if err := sqlDb.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	sqlDb.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	sqlDb.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Minute)

	log.Printf("Database connection established: %s", cfg.Postgres.DbName)
	return nil
}

func GetDb() *gorm.DB {
	return dbClient
}

func CloseDb() {
	connection, _ := dbClient.DB()
	connection.Close()
}
