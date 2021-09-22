package models

type Hai struct {
	Name       string    `form:"hai" binding:"required" json:"hai"`
}


type User struct {
	ID        int    `gorm:"primaryKey"`
	Email     string `gorm:"type:varchar(255)"`
	Password string  `gorm:"type:varchar(255)"`
}

