package main

import (
	"database/sql"

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

func getDb() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:IntelliP24@tcp(34.126.76.42:3306)/numberdb")
	if err != nil {
		panic(err.Error())
	}

	return db
}

var mainHandler = func(c *gin.Context) {
	phoneNumber := c.Param("phonenumber")
	pairs := phone.NewPairs(phoneNumber)
	bootPairs(pairs)

	db := getDb()
	repo := repositories.NewNumberRepo(db)
	bootRepoNumber(repo, pairs)

	c.JSON(200, gin.H{
		"API": API{
			PairsA: pairs.PairsA,
			Number: repo.PairNumberObjs,
		},
	})
}

func bootRepoNumber(repo *repositories.NumberRepo, pairs *phone.Pairs) {
	repo.SetTableNumbers(pairs.PairsUnique)
}

func bootPairs(pairs *phone.Pairs) {
	pairs.SetPairsA()
	pairs.SetPairsB()
	pairs.SetSumPhonenum()
	pairs.SetLastPair()
	pairs.SetPairsUnique()
}

type API struct {
	PairsA []string
	Number []repositories.PairNumberObj
}
