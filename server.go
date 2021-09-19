package main

import (
	"log"
	"net/http"
	"server/handlers"
	"server/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        int `gorm:"primaryKey"`
	Email     string
	CreatedAt time.Time
}

func main() {

	dsn := "root:lohith@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	router.GET("/", handlers.Home(router))

	router.GET("/:name", func(c *gin.Context) {
		c.String(http.StatusCreated, "w")
	})

	router.POST("/", func(c *gin.Context) {
		var hai models.Hai

		if err := c.ShouldBindJSON(&hai); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": string(hai.Name + "a"),
			"time":    string(hai.CreateTime.Hour()),
		})

	})

	router.GET("/dbtest/:id", func(c *gin.Context) {
		first := User{2, "lohith", time.Now()}
		result := db.Create(&first)
		println(result.Error)
	})

	http.ListenAndServe(":8080", router)
}
