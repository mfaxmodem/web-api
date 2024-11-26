package models

type User struct {
	BaseModel
	Username     string `gorm:"type:string;size:20;not null;unique"`
	FirstName    string `gorm:"type:string;size:15;default:null"` // null اگر میخواهید مقدار خالی باشد
	LastName     string `gorm:"type:string;size:25;default:null"` // null اگر میخواهید مقدار خالی باشد
	MobileNumber string `gorm:"type:string;size:11;unique;default:null"`
	Email        string `gorm:"type:string;size:64;unique;default:null"`
	Password     string `gorm:"type:string;size:64;not null"`
	Enabled      bool   `gorm:"default:true"`
	UserRoles    *[]UserRole
}
