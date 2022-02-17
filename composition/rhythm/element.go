package rhythm

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/notation/rat"
)

type Element struct {
	Duration rat.Rat
	Rest     bool
}

func NewElement(dur rat.Rat, rest bool) *Element {
	return &Element{
		Duration: dur,
		Rest:     rest,
	}
}

func (e *Element) String() string {
	durStr := e.Duration.String()
	if e.Rest {
		return fmt.Sprintf("(%s)", durStr)
	}

	return durStr
}

func (e *Element) Divide(n uint64) Elements {
	if n == 0 {
		return Elements{}
	}

	subdur := e.Duration.Div(rat.FromUint64(n))
	elems := make(Elements, n)

	for i := uint64(0); i < n; i++ {
		elems[i] = &Element{Duration: subdur, Rest: e.Rest}
	}

	return elems
}
