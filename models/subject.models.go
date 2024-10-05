package models

type Subject struct {
	ID          uint         `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string       `json:"name"`
	Code        string       `json:"code"`
	Credit      int          `json:"credit"`
	Sem         string       `json:"sem"`
	Branch      string       `json:"branch"`
	IsLab       bool         `json:"isLab" gorm:"default:false"`
	IsCore      bool         `json:"isCore" gorm:"default:false"`
	IsElective  bool         `json:"isElective" gorm:"default:false"`
	// Foreign key relation with Attendance
	Attendances []Attendance `json:"attendances" gorm:"foreignKey:SubjectID;constraint:OnDelete:CASCADE"`
}
