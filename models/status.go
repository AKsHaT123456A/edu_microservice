package models

type Stastus struct {
	ID        uint  `json:"statusId" gorm:"primaryKey;autoIncrement"`
	IsAc      bool  `json:"isAc" gorm:"default:false"`
	IsPresent bool  `json:"isPresent" gorm:"default:false"`
	IsAbsent  bool  `json:"isAbsent" gorm:"default:true"`
}