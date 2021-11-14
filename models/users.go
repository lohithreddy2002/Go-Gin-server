package models

import "gorm.io/gorm"

type Hai struct {
	Name       string    `form:"hai" binding:"required" json:"hai"`
}

type UserVerify struct {
	Id uint `gorm:"primaryKey type:autoIncrement"` 
	Email string `gorm:"type:varchar(255)" binding:"required" form:"email"`
	Otp string `gorm:"type:varchar(6)" binding:"required" form:"otp"`
}


type User struct {
	gorm.Model
	Email     string `gorm:"type:varchar(255)" binding:"required" form:"email"`
	Password string  `gorm:"type:varchar(255)" binding:"required" form:"password"`
}

