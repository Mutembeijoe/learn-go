package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Server-Time": time.Now().UTC(),
		})
	})

	r.Run()
}
