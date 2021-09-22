package models

import "gorm.io/gorm"

type Hai struct {
	Name       string    `form:"hai" binding:"required" json:"hai"`
}


type User struct {
	gorm.Model
	Email     string `gorm:"type:varchar(255)" binding:"required" form:"email"`
	Password string  `gorm:"type:varchar(255)" binding:"required" form:"password"`
}

