package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/test", mainHandler)

	r.Run(":3000")
}

var mainHandler = func(c *gin.Context) {
	c.JSON(200, gin.H{
		"Hello": "google",
	})
}
