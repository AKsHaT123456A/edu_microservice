package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"edumarshal.com/api/models"
	"edumarshal.com/api/utils"
)

// SubjectHandler handles subject creation requests.
func SubjectHandler(w http.ResponseWriter, r *http.Request) {
	var subject models.Subject
	if err := json.NewDecoder(r.Body).Decode(&subject); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := utils.CreateSubject(subject); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Subject created successfully"))
}

func AttendanceHandler(w http.ResponseWriter, r *http.Request) {
	var attendance models.Attendance
	if err := json.NewDecoder(r.Body).Decode(&attendance); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := utils.CreateAttendance(attendance); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Attendance created successfully"))
}

type Status struct {
	Date   string `json:"date"`   // Add JSON tags for marshaling
	Status string `json:"status"` // Add JSON tags for marshaling
}

func GetAttendanceHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	userIdStr := queryParams.Get("userId")
	subjectIdStr := queryParams.Get("subjectId")
	month := queryParams.Get("month")
	date := queryParams.Get("date")
	log.Println("userId:", userIdStr)
	log.Println("subjectId:", subjectIdStr)

	// Convert userId and subjectId to uint
	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid userId", http.StatusBadRequest)
		return
	}
	subjectId, err := strconv.ParseUint(subjectIdStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid subjectId", http.StatusBadRequest)
		return
	}

	// Cast to uint
	userIdUint := uint(userId)
	subjectIdUint := uint(subjectId)

	// Get the attendance records
	attendanceRetrieved, err := utils.GetUserAttendance(userIdUint, subjectIdUint, date,month)
	if err != nil {
		http.Error(w, "Failed to retrieve attendance records: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Collect statuses
	var statuses []Status
	for _, attendance := range attendanceRetrieved {
		statuses = append(statuses, Status{
			Date:   attendance.Date,   // Format the date as needed
			Status: attendance.Status, // Convert attendance.Status to a slice of strings
		})
	}

	// Set response content type to application/json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Marshal the statuses array to JSON
	statusJSON, err := json.Marshal(statuses)
	if err != nil {
		log.Println("Error converting status to JSON:", err)
		http.Error(w, "Error converting status to JSON", http.StatusInternalServerError)
		return
	}

	// Write the JSON response
	w.Write(statusJSON)
}
