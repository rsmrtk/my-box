package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Congratulations!",
		})
	})

	if err := r.Run(`:9595`); err != nil {
		panic(err)
	}
}
