package repositories

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type NumberRepo struct {
	Db *sql.DB
}

type PairNumber struct {
	PairNumber string
	PairType   string
	PairPoint  int
}

func NewNumberRepo(db *sql.DB) *NumberRepo {
	return &NumberRepo{
		Db: db,
	}
}

func (numberRepo *NumberRepo) GetTableNumbers() []PairNumber {
	var pairNumber PairNumber
	var pairNumbers []PairNumber
	rows, _ := numberRepo.Db.Query("SELECT pairnumber, pairtype, pairpoint FROM numbers")
	for rows.Next() {
		err := rows.Scan(&pairNumber.PairNumber, &pairNumber.PairType, &pairNumber.PairPoint)
		if err != nil {
			panic(err)
		}

		if err = rows.Err(); err != nil {
			panic(err)
		}

		number := PairNumber{PairNumber: pairNumber.PairNumber, PairType: pairNumber.PairType, PairPoint: pairNumber.PairPoint}
		pairNumbers = append(pairNumbers, number)
	}
	return pairNumbers
}
