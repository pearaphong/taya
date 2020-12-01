package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pearaphong/taya/phone"
)

func main() {
	r := gin.Default()

	r.GET("/pairs/:phonenumber", mainHandler)

	r.Run(":3000")
}

var mainHandler = func(c *gin.Context) {
	phoneNumber := c.Param("phonenumber")
	pairs := phone.NewPairs(phoneNumber)
	bootPairs(pairs)

	c.JSON(200, gin.H{
		"PairsA": pairs.PairsA,
		"PairsB": pairs.PairsB,
	})
}

func bootPairs(pairs *phone.Pairs) {
	pairs.SetPairsA()
	pairs.SetPairsB()
}
