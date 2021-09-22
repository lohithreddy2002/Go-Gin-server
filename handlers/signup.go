package handlers

import (
	"net/http"
	"server/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Signup(router *gin.Engine,db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		err := c.ShouldBind(&user)

		if err == nil{
			println(user.Password)
			db.Create(&user)
			c.JSON(http.StatusOK,gin.H{
				"message":"success",
			})
		} else {
			c.JSON(http.StatusBadRequest,gin.H{
				"message" : "failure",
			})
		}

	}
}