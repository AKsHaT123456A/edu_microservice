package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	config "edumarshal.com/api/config"
	"edumarshal.com/api/models"
)

// DB establishes a connection to the PostgreSQL database using a DSN and returns the DB instance.
func DB() (*gorm.DB, error) {
	// Your PostgreSQL connection string
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
		return nil, err
	}
	dsn:=config.Link
	// Open a connection to the PostgreSQL database using Gorm
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // Disable FK constraints when migrating
	})
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to the database!")

	// Auto-migrate the models to keep the database schema up to date
	err = db.AutoMigrate(
		&models.User{},
		&models.Attendance{},
		&models.Status{},
	)
	if err != nil {
		return nil, err
	}

	log.Println("Database migration completed successfully!")
	return db, nil
}
