package meter

import "github.com/jamestunnell/go-musicality/notation/rat"

func FourFour() *Meter {
	return New(4, rat.New(1, 4))
}

func ThreeFour() *Meter {
	return New(3, rat.New(1, 4))
}

func TwoFour() *Meter {
	return New(2, rat.New(1, 4))
}

func TwoTwo() *Meter {
	return New(2, rat.New(1, 2))
}

func SixEight() *Meter {
	return New(6, rat.New(3, 8))
}
