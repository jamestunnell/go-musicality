package meter

import "math/big"

func FourFour() *Meter {
	return New(4, big.NewRat(1, 4))
}

func ThreeFour() *Meter {
	return New(3, big.NewRat(1, 4))
}

func TwoFour() *Meter {
	return New(2, big.NewRat(1, 4))
}

func TwoTwo() *Meter {
	return New(2, big.NewRat(1, 2))
}

func SixEight() *Meter {
	return New(2, big.NewRat(3, 8))
}
