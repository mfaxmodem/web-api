package models

type User struct {
	BaseModel
	Username     string  `gorm:"type:varchar(20);not null;unique"`
	FirstName    *string `gorm:"type:varchar(15)"` // استفاده از اشاره‌گر برای پذیرش NULL
	LastName     *string `gorm:"type:varchar(25)"`
	MobileNumber *string `gorm:"type:varchar(11);unique"`
	Email        *string `gorm:"type:varchar(64);unique"`
	Password     string  `gorm:"type:varchar(64);not null"`
	Enabled      bool    `gorm:"default:true"`
	UserRoles    *[]UserRole
}
