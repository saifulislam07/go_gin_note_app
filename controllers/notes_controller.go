package controllers

import (
	"hello_gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotesIndex(c *gin.Context) {
	notes := models.NotesAll()

	c.HTML(http.StatusOK, "notes/index.html",
		gin.H{
			"notes": notes,
		},
	)
}

func NotesNew(c *gin.Context) {
	c.HTML(http.StatusOK, "notes/new.html",
		gin.H{},
	)
}

func NotesCreate(c *gin.Context) {
	name := c.PostForm("name")
	content := c.PostForm("content")

	models.NotesCreate(name, content)

	c.Redirect(http.StatusSeeOther, "/notes")
}
