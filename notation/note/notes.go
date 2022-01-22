package note

import "math/big"

type Notes []*Note

func (notes Notes) TotalDuration() *big.Rat {
	sum := big.NewRat(0, 1)

	for _, n := range notes {
		sum.Add(sum, n.Duration)
	}

	return sum
}
