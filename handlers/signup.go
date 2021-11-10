package handlers

import (
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Signup(router *gin.Engine, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "failure",
			})
			return
		}
		err = db.Where("Email = ?", user.Email).First(&user).Error
		if err == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "failure",
			})
			return
		}
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(hashedPassword)

		result := db.Create(&user)

		if result.Error != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "Error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})

	}
}
