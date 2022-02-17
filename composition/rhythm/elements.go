package rhythm

import "github.com/jamestunnell/go-musicality/notation/rat"

type Elements []*Element

func (elems Elements) Strings() []string {
	elemStrings := make([]string, len(elems))

	for i, e := range elems {
		elemStrings[i] = e.String()
	}

	return elemStrings
}

func (elems Elements) Duration() rat.Rat {
	sum := rat.Zero()

	for _, e := range elems {
		sum = sum.Add(e.Duration)
	}

	return sum
}
