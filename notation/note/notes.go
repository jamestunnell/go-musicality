package note

import "github.com/jamestunnell/go-musicality/notation/rat"

type Notes []*Note

func (notes Notes) TotalDuration() rat.Rat {
	sum := rat.Zero()

	for _, n := range notes {
		sum = sum.Add(n.Duration)
	}

	return sum
}
