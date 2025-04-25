package main

import (
	"hello_gin/controllers"
	"hello_gin/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	r.Static("/static", "./static")

	r.LoadHTMLGlob("templates/**/**")

	models.ConnectDatabase()
	models.DBMigrate()

	r.GET("/notes", controllers.NotesIndex)
	r.GET("/notes/new", controllers.NotesNew)
	r.POST("/notes", controllers.NotesCreate)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/index.html", gin.H{
			"title":   "Notes application",
			"message": "Welcome to your notes app!",
		})
	})

	log.Println("🚀 Server started at http://localhost:8090")
	r.Run(":8090")
}
