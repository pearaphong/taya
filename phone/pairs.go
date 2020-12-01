package phone

type Pairs struct {
	phoneNumber string
	PairsA      []string
	PairsB      []string
}

func NewPairs(phoneNumber string) *Pairs {
	return &Pairs{
		phoneNumber: phoneNumber,
	}
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
