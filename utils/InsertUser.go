package utils

import (
	"fmt"
	"log"

	"edumarshal.com/api/db"
	"edumarshal.com/api/models"
)

// InsertUser inserts a new user into the database
func InsertUser(user models.User) error {
	db, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database connection: %w", err)
	}

	// Create the user record
	if err := db.Create(&user).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	log.Println("User created successfully")
	return nil
}

// GetUserByUsername retrieves a user by username from the database
func GetUserByUsername(username string) (*models.User, error) {
	db, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database connection: %w", err)
	}

	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}
	fmt.Println(user)
	return &user, nil
}
