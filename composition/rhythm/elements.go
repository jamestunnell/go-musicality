package rhythm

import "github.com/jamestunnell/go-musicality/notation/rat"

type Elements []*Element

func (elems Elements) Duration() rat.Rat {
	sum := rat.Zero()

	for _, e := range elems {
		sum = sum.Add(e.Duration)
	}

	return sum
}
