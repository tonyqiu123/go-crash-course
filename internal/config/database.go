package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDatabase initializes the database connection
func InitDatabase(cfg *Config) *gorm.DB {
	// Configure GORM logger based on environment
	var logLevel logger.LogLevel
	if cfg.Environment == "production" {
		logLevel = logger.Error
	} else {
		logLevel = logger.Info
	}

	// Open database connection
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{
		Logger: logger.Default.LogLevel(logLevel),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate models
	// TODO: Call AutoMigrate for all models
	// err = db.AutoMigrate(
	// 	&clubs.Clubs{},
	// 	&events.Events{},
	// 	&events.EventDates{},
	// 	&events.EventSubmission{},
	// 	&events.EventInterest{},
	// 	&events.IgnoredPost{},
	// 	// ... other models
	// )
	// if err != nil {
	// 	log.Fatalf("Failed to auto-migrate: %v", err)
	// }

	log.Println("Database connection established")
	return db
}
