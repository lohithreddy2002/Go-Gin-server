package handlers

import
(
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"server/models"
	"net/http"
)

func VerifyOtp(router *gin.Engine, db *gorm.DB) gin.HandlerFunc{

	var input models.UserVerify
	return func(c * gin.Context){
		err:= c.ShouldBind(&input)
		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":"send correct fields",
			})
			return
		}
		var otpmodel models.UserVerify
		err = db.Where("email = ?",input.Email).Last(&otpmodel).Error

		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"otp":"Error",
			})
			return
		}
		
		if(otpmodel.Otp != input.Otp){
			c.JSON(http.StatusBadRequest,gin.H{
				"otp":"Error",
			})
			return
		}
		var user models.User
		err = db.Where("email = ?",otpmodel.Email).First(&user).Error

		if err != nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"otp":"Error",
			})
			return
		}
		db.Where("email = ?",otpmodel.Email).Delete(models.UserVerify{})
		token,err := CreateToken(user)
		if err != nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"otp":"Error",
			})
			return
		}

 		c.JSON(http.StatusOK,gin.H{
			"token":token,
		})
	}

}