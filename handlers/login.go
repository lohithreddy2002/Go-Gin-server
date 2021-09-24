package handlers

import (
	"net/http"
	"server/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

func Login(router *gin.Engine,db *gorm.DB) gin.HandlerFunc {
	var input models.User
	return func(c *gin.Context) {
		err := c.ShouldBind(&input)
		if err != nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"message" : "bad error",
			})
			return
		}
		var user models.User
		err = db.Where("email = ?",input.Email).First(&user).Error
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"message" : "user not found",
			})
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
			c.JSON(http.StatusBadRequest,gin.H{
				"message" : "password not same",
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"message" : "user loged in",
		})
			
}
}

