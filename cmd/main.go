package main

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})
	r.Static("/assets", "frontend/assets")
	r.LoadHTMLGlob("frontend/templates/*.html")

	// Web pages
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": "This is a content coming from golang...",
		})
	})

	// API endpoints
	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "this is coming from an API call",
		})
	})

	r.Run("localhost:8080")
}
