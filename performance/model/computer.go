package model

import (
	"fmt"
	"math/big"

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

// NewComputer takes an initial value and a set of changes and
// produces a Computer which can calculate the value at any time.
// Assumes that the changes are valid.
func NewComputer(startVal float64, changes change.Map) (*Computer, error) {
	// if result := changes.Validate(); result != nil {
	// 	return nil, result
	// }

	pairs := []function.SubdomainFunctionPair{}

	makeComputer := func() (*Computer, error) {
		pwf, err := function.NewPiecewiseFunction(pairs)
		return &Computer{pwf}, err
	}

	if len(changes) == 0 {
		pairs = append(pairs, function.SubdomainFunctionPair{
			Subdomain: function.DomainAll(),
			Function:  function.NewConstantFunction(startVal),
		})

		return makeComputer()
	}

	offsets := changes.SortedOffsets()

	if err := checkChangeOverlap(changes, offsets); err != nil {
		return nil, err
	}

	n := len(offsets)
	prevChangeEnd := function.DomainMin()
	prevEndVal := startVal

	for i := 0; i < n; i++ {
		offset := offsets[i]

		// if there is a gap, fill in with a constant function
		if offset.Cmp(prevChangeEnd) == 1 {
			pairs = append(pairs, function.SubdomainFunctionPair{
				Subdomain: function.NewRange(prevChangeEnd, offset),
				Function:  function.NewConstantFunction(prevEndVal),
			})
		}

		change := changes[offset]

		if change.Duration.Cmp(zero) == 0 {
			// Don't do anything here. A constant function will be added at the
			// beginning of next loop, or just after the loop ends if it's the last change.

			prevChangeEnd = offset
		} else {
			end := new(big.Rat).Add(offset, change.Duration)

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

func checkChangeOverlap(changes change.Map, offsets change.Rats) error {
	n := len(changes)

	for i := 0; i < (n - 1); i++ {
		curr := changes[offsets[i]]
		end := new(big.Rat).Add(offsets[i], curr.Duration)

		if end.Cmp(offsets[i+1]) == 1 {
			return fmt.Errorf("change at offset %v overlaps next", offsets[i])
		}
	}

	return nil
}
