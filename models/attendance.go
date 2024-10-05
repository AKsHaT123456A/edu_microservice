package models

type Attendance struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint   `json:"userId"`    // Foreign key to User
	SubjectID uint   `json:"subjectId"` // Foreign key to Subject
	Date      string `json:"date"`
	Status    string `json:"status"`
	Month     string `json:"month"`
	// Relationships
	User    User    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Subject Subject `gorm:"foreignKey:SubjectID;constraint:OnDelete:CASCADE"`
}
