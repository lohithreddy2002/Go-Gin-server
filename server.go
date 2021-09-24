package main

import (
	"log"
	"net/http"
	"server/handlers"
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

	router.GET("/", handlers.Home(router))
	
	router.POST("/signup",handlers.Signup(router,db))

	router.POST("/login",handlers.Login(router,db))

	http.ListenAndServe(":8080", router)
}
