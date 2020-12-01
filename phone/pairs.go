package phone

import "strconv"

type Pairs struct {
	phoneNumber  string
	PairsA       []string
	PairsB       []string
	PairSumPhone string
	PairLastA    string
	PairsUnique  []string
}

func NewPairs(phoneNumber string) *Pairs {
	return &Pairs{
		phoneNumber: phoneNumber,
	}
}

func (pairs *Pairs) SetPairsUnique() {
	var mixPairsAB []string
	for _, pair := range pairs.PairsA {
		mixPairsAB = append(mixPairsAB, pair)
	}

	for _, pair := range pairs.PairsB {
		mixPairsAB = append(mixPairsAB, pair)
	}
	pairs.PairsUnique = unique(mixPairsAB)
}

func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func (pairs *Pairs) SetLastPair() {
	var lastPair string
	numberRune := []rune(pairs.phoneNumber)
	lastPair = string(numberRune[8:10])
	pairs.PairLastA = lastPair
}

func (pairs *Pairs) SetSumPhonenum() {
	var sum int

	numberRune := []rune(pairs.phoneNumber)

	for _, c := range numberRune {
		num, _ := strconv.Atoi(string(c))
		sum = sum + num
	}

	pairs.PairSumPhone = strconv.Itoa(sum)

}

func (pairs *Pairs) SetPairsA() {

	var pairsA []string
	numberRune := []rune(pairs.phoneNumber)
	pairsA = append(pairsA, string(numberRune[0:2]))
	pairsA = append(pairsA, string(numberRune[2:4]))
	pairsA = append(pairsA, string(numberRune[4:6]))
	pairsA = append(pairsA, string(numberRune[6:8]))
	pairsA = append(pairsA, string(numberRune[8:10]))
	pairs.PairsA = pairsA

}

func (pairs *Pairs) SetPairsB() {

	var pairsB []string
	numberRune := []rune(pairs.phoneNumber)
	pairsB = append(pairsB, string(numberRune[0:1]))
	pairsB = append(pairsB, string(numberRune[1:3]))
	pairsB = append(pairsB, string(numberRune[3:5]))
	pairsB = append(pairsB, string(numberRune[5:7]))
	pairsB = append(pairsB, string(numberRune[7:9]))
	pairsB = append(pairsB, string(numberRune[9:11]))
	pairs.PairsB = pairsB

}
