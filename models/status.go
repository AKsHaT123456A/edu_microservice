package models

type Status struct {
	StatusId     int64 `json:"statusId" gorm:"column:StatusId;primaryKey;autoIncrement"`
	IsAc         bool  `json:"isAc" gorm:"default:false"`
	IsPresent    bool  `json:"isPresent" gorm:"default:false"`
	IsAbsent     bool  `json:"isAbsent" gorm:"default:true"`
	AttendanceID uint  `json:"attendanceId"` // Foreign key
}
