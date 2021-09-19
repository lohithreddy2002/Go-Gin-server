package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(router *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		router.LoadHTMLGlob("templates/*")
		c.HTML(http.StatusOK, "1.html", gin.H{"title": "hai"})
	}
}
