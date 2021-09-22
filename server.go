package main

import (
	"log"
	"net/http"
	"server/handlers"
	"server/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
)

func main() {

	router := gin.Default()
	dsn := "root:lohith@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		log.Fatal(err)
		println("ERROR")
	}

	router.GET("/", handlers.Home(router))

	router.GET("/:name", func(c *gin.Context) {
		var test models.User
		result := db.First(&test)
		c.JSON(http.StatusOK, gin.H{
			"user": test.Email,
		})
		fmt.Println(result.Error)

	})

	router.POST("/", func(c *gin.Context) {
		// var hai models.Hai
		// err := c.ShouldBindJSON(&hai)
		message := c.PostForm("hai")
		test := c.DefaultPostForm("test","test")
		c.JSON(http.StatusOK,gin.H{"message":message,"test":test})
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// }

	})

	router.GET("/dbtest/:id", func(c *gin.Context) {
		db.AutoMigrate(&models.User{})
		first := models.User{4, "lohith","test"}
		result := db.Create(&first)
		fmt.Println(result.RowsAffected)
		c.JSON(http.StatusOK,gin.H{"message":string(first.ID)})
	})

	http.ListenAndServe(":8080", router)
}
