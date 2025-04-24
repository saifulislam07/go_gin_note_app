package main

import (
	"hello_gin/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectDatabase() {
	dsn := "gin_notes:password@tcp(127.0.0.1:3306)/gin_notes?parseTime=true"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("❌ Failed to connect to database: " + err.Error())
	}

	log.Println("✅ GORM database connection successful!")
}

func dbMigrate() {
	err := DB.AutoMigrate(&models.Note{})
	if err != nil {
		panic("❌ Migration failed: " + err.Error())
	}
	log.Println("✅ Database migration complete!")
}

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	r.Static("/vanilla", "./static/vendor")
	r.LoadHTMLGlob("templates/**/**")

	connectDatabase()
	dbMigrate()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/index.html", gin.H{
			"title":   "Notes application",
			"message": "Welcome to your notes app!",
		})
	})

	log.Println("🚀 Server started at http://localhost:8090")
	r.Run(":8090")
}
