package handlers

import (
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/smtp"
	"fmt"
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

func Sendemail(email string,otp string) error{
	from:="fastgram722@gmail.com"
	pswd:= "Fastgram23"
	host := "smtp.gmail.com"
 
    // Its the default port of smtp server
    port := "587"

	println(email)
 
	toList := []string{"rocklohithreddy@gmail.com"}

    // This is the message to send in the mail
    msg := "Hello User your otp is "+otp

	body:= []byte(msg)

	auth:= smtp.PlainAuth("",from,pswd,host)
	err := smtp.SendMail(host+":"+port, auth, from, toList, body)
 
    // handling the errors
    if err != nil {
        fmt.Println(err)
		return err
    }
 
    fmt.Println("Successfully sent mail to all user in toList")

	return nil



}
