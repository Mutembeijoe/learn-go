package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mutembeijoe/data_io/project_1_url_shortner/base62"
	"github.com/mutembeijoe/data_io/project_1_url_shortner/database"
	"net/http"
)

func (app *application) generateShortUrl(c *gin.Context) {
	var url database.UrlResource
	var urlId int

	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	err := app.DB.QueryRow("INSERT INTO urls(originalurl) VALUES ($1) RETURNING id", url.OriginalUrl).Scan(&urlId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server error",
		})
	}

	shortenedUrl := base62.ToBase62(urlId)

	c.JSON(http.StatusOK, gin.H{
		"url": shortenedUrl,
	})
}
