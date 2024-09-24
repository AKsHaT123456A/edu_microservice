package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	config "edumarshal.com/api/config"
	models "edumarshal.com/api/models"
)

func DB() (*gorm.DB, error) {
	config, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
		return nil, err
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Successfully connected to the database!")

	if err := db.AutoMigrate(&models.User{}, &models.Attendance{}, &models.Status{}); err != nil {
		return nil, fmt.Errorf("failed to migrate models: %w", err)
	}

	return db, nil
}
