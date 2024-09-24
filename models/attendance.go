package models

type Attendance struct {
	ID     uint     `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID uint     `json:"userId"` 
	Date   string   `json:"date"`
	Status []Status `json:"status" gorm:"foreignKey:AttendanceID"`
}
