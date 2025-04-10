package db

import (
	"fmt"

	"github.com/anujsinghrawat/event-manager/config"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(config *config.EnvConfig, DBMigrator func(db *gorm.DB) error) *gorm.DB {
	uri := fmt.Sprintf(`
		host=%s user=%s password=%s dbname=%s sslmode=%s port=5432`,
		config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBSSLMode,
	)
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Unable to connect to database: %e", err)
	}

	log.Info("Connected to database!")

	if err := DBMigrator(db); err != nil {
		log.Fatalf("Error migrating database: %e", err)
	}
	return db
}
