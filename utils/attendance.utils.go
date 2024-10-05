package utils

import (
	"log"

	"edumarshal.com/api/db"
	"edumarshal.com/api/models"
)

// CreateAttendance creates a new attendance record in the database.
func CreateAttendance(attendance models.Attendance) error {
	database, err := db.DB()
	if err != nil {
		log.Printf("Failed to get database connection: %v", err)
		return err
	}

	if err := database.Create(&attendance).Error; err != nil {
		log.Printf("Failed to create attendance: %v", err)
		return err
	}
	log.Println("Attendance created successfully")
	return nil
}

func GetUserAttendance(userID uint, subjectID uint, date string, month string) ([]models.Attendance, error) {
	var attendances []models.Attendance
	log.Println(userID, subjectID, date, month)
	// Get the GORM instance from db
	database, err := db.DB() // Assuming db.DB() returns *gorm.DB
	if err != nil {
		log.Printf("Failed to get database connection: %v", err)
		return nil, err
	}

	// Start building the query with userID and subjectID conditions
	query := database.Preload("User").Preload("Subject").
		Where("user_id = ? AND subject_id = ?", userID, subjectID)
	
	if month != "" &&date=="" {
		log.Println("Query:", month)
		query = database.
			Where("user_id = ? AND subject_id = ? and month= ?", userID, subjectID, month)
	}
	// If a date is provided (not empty), add the date condition to the query
	if date != "" && month=="" {
		// Add the date filter to the query
		query = database.
			Where("user_id = ? AND subject_id = ? and date= ?", userID, subjectID, date)
	}

	sql, args := query.Statement.SQL.String(), query.Statement.Vars
	log.Printf("Query: %s, Args: %v", sql, args)
	// Execute the query to find attendance records
	result := query.Find(&attendances)
	if result.Error != nil {
		log.Printf("Failed to retrieve attendance records: %v", result.Error)
		return nil, result.Error
	}

	return attendances, nil
}
