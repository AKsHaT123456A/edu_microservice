package models

type User struct {
	ID           uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	ProfileImage string `json:"profileImage"`
	Username     string `json:"username"`
    StudentName string `json:"studentName"`
    Email        string `json:"email"`
    StudentNumber string `json:"studentNumber"`
    UniversityRollNumber int64 `json:"universityRollNumber"`
	PasswordHash string `json:"passwordHash"`
	DOB          string `json:"dob"`
	IsRemembered bool   `json:"isRemembered" gorm:"default:false"`
}
