package repositories

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type NumberRepo struct {
	Db             *sql.DB
	PairNumberObjs []PairNumberObj
}

type PairNumberObj struct {
	PairNumber string
	PairType   string
	PairPoint  int
}

func NewNumberRepo(db *sql.DB) *NumberRepo {
	return &NumberRepo{
		Db: db,
	}
}

func (numberRepo *NumberRepo) SetTableNumbers(uniqueNumber []string) {
	var pairNumber PairNumberObj
	var pairNumbers []PairNumberObj

	rows, _ := numberRepo.Db.Query("SELECT pairnumber, pairtype, pairpoint FROM numbers")
	defer numberRepo.Db.Close()
	for rows.Next() {
		err := rows.Scan(&pairNumber.PairNumber, &pairNumber.PairType, &pairNumber.PairPoint)
		if err != nil {
			panic(err)
		}

		if err = rows.Err(); err != nil {
			panic(err)
		}

		number := PairNumberObj{PairNumber: pairNumber.PairNumber, PairType: pairNumber.PairType, PairPoint: pairNumber.PairPoint}
		pairNumbers = append(pairNumbers, number)

	}

	for _, unique := range uniqueNumber {
		for _, pair := range pairNumbers {
			if unique == pair.PairNumber {
				numberRepo.PairNumberObjs = append(numberRepo.PairNumberObjs, pair)
				break
			}
		}
	}

}
