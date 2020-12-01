package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
	"github.com/pearaphong/taya/phone"
	"github.com/pearaphong/taya/repositories"
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

	db, err := sql.Open("mysql", "root:IntelliP24@tcp(34.126.76.42:3306)/numberdb")
	defer db.Close()
	if err != nil {
		log.Fatal("connect fail")
	} else {
		log.Print("connected")
	}
	numberRepo := repositories.NewNumberRepo(db)

	c.JSON(200, gin.H{
		"API": API{Pairs: *pairs, Number: numberRepo.GetTableNumbers()},
	})
}

func bootPairs(pairs *phone.Pairs) {
	pairs.SetPairsA()
	pairs.SetPairsB()
	pairs.SetSumPhonenum()
	pairs.SetLastPair()
	pairs.SetPairsUnique()
}

type API struct {
	Pairs  phone.Pairs
	Number []repositories.PairNumber
}
