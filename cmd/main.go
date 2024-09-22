package main

import "github.com/gin-gonic/gin"

const port = ":8080"

func main() {
	r := gin.Default()
	r.GET("/games", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Game list"})
	})
	r.Run(port)
}
