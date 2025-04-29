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

	r.GET("/login", controllers.LoginPage)
	r.GET("/signup", controllers.SignupPage)

	r.POST("/login", controllers.Login)
	r.POST("/signup", controllers.SignupCreate)
	r.POST("/logout", controllers.Logout)

	r.GET("/notes", controllers.NotesIndex)
	r.GET("/notes/new", controllers.NotesNew)
	r.POST("/notes", controllers.NotesCreate)
	r.GET("/notes/:id", controllers.NotesShow)
	r.GET("/notes/edit/:id", controllers.NotesEdit)
	r.POST("/notes/edit/:id", controllers.NotesUpdate)
	r.POST("/notes/:id/delete", controllers.NotesDelete)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.html", gin.H{
			"title":   "Notes application",
			"message": "Welcome to your notes app!",
		})
	})

	log.Println("🚀 Server started at http://localhost:8090")
	r.Run(":8090")
}
