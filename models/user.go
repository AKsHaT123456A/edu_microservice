package models

type User struct {
	ID                   uint         `json:"id" gorm:"primaryKey;autoIncrement"`
	ProfileImage         string       `json:"profileImage"`
	Username             string       `json:"username"`
	StudentName          string       `json:"studentName"`
	Email                string       `json:"email"`
	StudentNumber        string       `json:"studentNumber"`
	UniversityRollNumber int64        `json:"universityRollNumber"`
	PasswordHash         string       `json:"passwordHash"`
	DOB                  string       `json:"dob"`
	Branch               string       `json:"branch"`
	Hostel               string       `json:"hostel" gorm:"default:'Not Alloted'"`
	Semester             string       `json:"semester" gorm:"default:'Even'"`
	AdmissionDate        string       `json:"admissionDate"`
	AdmissionMode        string       `json:"admissionMode" gorm:"default:'JEE'"`
	Categories           string       `json:"categories" gorm:"default:'General'"`
	JeeRank              string       `json:"jeeRank" gorm:"default:'Not Specified'"`
	IsLateral            bool         `json:"isLateral" gorm:"default:false"`
	Section              string       `json:"section"`
	Course               string       `json:"course"`
	IsRemembered         bool         `json:"isRemembered" gorm:"default:false"`
	IsAdmin              bool         `json:"isAdmin" gorm:"default:false"`
	// Foreign key relation with Attendance
	Attendances          []Attendance `json:"attendances" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}


type UserLoginRequest struct {
	Username     string `json:"username"`
	DOB          string `json:"dob"`
	PasswordHash string `json:"passwordHash"`
	Role         string `json:"role"`
}
