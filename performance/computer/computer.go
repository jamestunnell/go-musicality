package computer

import (
	"fmt"
	"sort"

	"github.com/jamestunnell/go-musicality/notation/change"
	"github.com/jamestunnell/go-musicality/performance/function"
)

// Computer represents an initial value and a set of changes as a
// piecewise function, calculating the current value at any time.
// Changes to the initial value can be immediate or gradual (see change functions).
// Between periods of transition, the value will be constant.
type Computer struct {
	*function.PiecewiseFunction
}

// New takes an initial value and a set of changes and
// produces a Computer which can calculate the value at any time.
// Assumes that the changes are valid.
func New(startVal float64, changes change.Changes) (*Computer, error) {
	sort.Sort(changes)

	changes = SimplifyChanges(startVal, changes)

	if err := checkChangeOverlap(changes); err != nil {
		return nil, err
	}

	pairs := []function.SubdomainFunctionPair{}

	makeComputer := func() (*Computer, error) {
		pwf, err := function.NewPiecewiseFunction(pairs)
		return &Computer{pwf}, err
	}

	if changes.Len() == 0 {
		pairs = append(pairs, function.SubdomainFunctionPair{
			Subdomain: function.DomainAll(),
			Function:  function.NewConstantFunction(startVal),
		})

		return makeComputer()
	}

	prevChangeEnd := function.DomainMin()
	prevEndVal := startVal

	for _, change := range changes {
		offset := change.Offset

		// if there is a gap, fill in with a constant function
		if offset.Greater(prevChangeEnd) {
			pairs = append(pairs, function.SubdomainFunctionPair{
				Subdomain: function.NewRange(prevChangeEnd, offset),
				Function:  function.NewConstantFunction(prevEndVal),
			})
		}

		if change.Duration.Zero() {
			// Don't do anything here. A constant function will be added at the
			// beginning of next loop, or just after the loop ends if it's the last change.

			prevChangeEnd = offset
		} else {
			end := offset.Add(change.Duration)

			p1 := function.NewPoint(offset, prevEndVal)
			p2 := function.NewPoint(end, change.EndValue)
			f := function.NewLinearFunctionFromPoints(p1, p2)

			pairs = append(pairs, function.SubdomainFunctionPair{
				Subdomain: function.NewRange(offset, end),
				Function:  f,
			})

			prevChangeEnd = end
		}

		prevEndVal = change.EndValue
	}

	// until the end of time
	pairs = append(pairs, function.SubdomainFunctionPair{
		Subdomain: function.NewRange(prevChangeEnd, function.DomainMax()),
		Function:  function.NewConstantFunction(prevEndVal),
	})

	return makeComputer()
}

func checkChangeOverlap(changes change.Changes) error {
	n := changes.Len()

	// check for change overlap
	for i := 0; i < (n - 1); i++ {
		curr := changes[i]
		end := changes[i].Offset.Add(curr.Duration)

		if end.Greater(changes[i+1].Offset) {
			return fmt.Errorf("change at offset %s overlaps next", curr.Offset.String())
		}
	}

	return nil
}
