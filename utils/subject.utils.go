package utils

import (
	"log"

	"edumarshal.com/api/db"
	"edumarshal.com/api/models"
)

// CreateSubject creates a new subject in the database.
func CreateSubject(subject models.Subject) error {
	database, err := db.DB()
	if err != nil {
		log.Printf("Failed to get database connection: %v", err)
		return err
	}

	if err := database.Create(&subject).Error; err != nil {
		log.Printf("Failed to create subject: %v", err)
		return err
	}
	log.Println("Subject created successfully")
	return nil
}
