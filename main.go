package main

import (
	"hello_gin/controllers"
	"hello_gin/middlewares"
	"hello_gin/models"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	r.Static("/static", "./static")

	r.LoadHTMLGlob("templates/**/**")

	models.ConnectDatabase()
	models.DBMigrate()

	store := memstore.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("notes", store))

	r.Use(middlewares.AuthenticateUser())

	r.GET("/login", controllers.LoginPage)
	r.GET("/signup", controllers.SignupPage)

	r.POST("/login", controllers.Login)
	r.POST("/signup", controllers.Signup)
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
			"title":     "Notes application",
			"logged_in": (c.GetUint64("user_id") > 0),
		})
	})

	log.Println("🚀 Server started at http://localhost:8090")
	r.Run(":8090")
}
