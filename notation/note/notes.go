package note

import (
	"math/big"

	"github.com/jamestunnell/go-musicality/common/rat"
)

type Notes []*Note

func (notes Notes) TotalDuration() *big.Rat {
	sum := rat.Zero()

	for _, n := range notes {
		sum = rat.Add(sum, n.Duration)
	}

	return sum
}
