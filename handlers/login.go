package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"server/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(router *gin.Engine, db *gorm.DB) gin.HandlerFunc {
	var input models.User
	return func(c *gin.Context) {
		err := c.ShouldBind(&input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "bad error",
			})
			return
		}
		var user models.User
		err = db.Where("email = ?", input.Email).First(&user).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "user not found",
			})
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "password not same",
			})
			return
		}

		token, err := CreateToken(user)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "user loged in",
			"token":   token,
		})

	}
}

var jwtKey = []byte("FDr1VjVQiSiybYJrQZNt8Vfd7bFEsKP6vNX1brOSiWl0mAIVCxJiR4/T3zpAlBKc2/9Lw2ac4IwMElGZkssfj3dqwa7CQC7IIB+nVxiM1c9yfowAZw4WQJ86RCUTXaXvRX8JoNYlgXcRrK3BK0E/fKCOY1+izInW3abf0jEeN40HJLkXG6MZnYdhzLnPgLL/TnIFTTAbbItxqWBtkz6FkZTG+dkDSXN7xNUxlg==")

type authClaims struct {
	jwt.StandardClaims
	UserID uint `json:"userId"`
}

func CreateToken(user models.User) (string, error) {
	expiresat := time.Now().Add(24 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, authClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   user.Email,
			ExpiresAt: expiresat,
		},
		UserID: user.ID,
	})
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func VerifyToken(tokenString string) (uint, string, error) {
	var claims authClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		return 0, "", err
	}
	if !token.Valid {
		return 0, "", errors.New("invalid token")
	}
	id := claims.UserID
	username := claims.Subject
	return id, username, nil
}
