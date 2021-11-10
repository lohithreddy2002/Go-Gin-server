package main

import (
	"net/http"
	"server/handlers"
	"server/models"
	"log"
	"server/middleware"
	"github.com/gin-gonic/gin"
		"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	router := gin.Default()
	dsn := "root:lohith@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		println("ERROR")
	}

	db.AutoMigrate(models.User{})

	router.GET("/", handlers.Home(router))

	authorized := router.Group("/v1/")

	authorized.Use(middleware.CheckToken)
	{
		authorized.GET("/test",func(c *gin.Context){
			id,err := c.Get("id")
			c.JSON(http.StatusOK,gin.H{
				"id":id,
				"err":err,
			})
		})
	
	}


	router.POST("/signup", handlers.Signup(router, db))

	router.POST("/login", handlers.Login(router, db))

	
	http.ListenAndServe(":8080", router)
}
