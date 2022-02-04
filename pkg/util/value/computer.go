package value

import (
	"fmt"
	"math"
	"sort"

	"github.com/jamestunnell/go-musicality/performance/function"
	"github.com/jamestunnell/go-musicality/pkg/util"
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
func NewComputer(startVal float64, changes map[float64]*Change) (*Computer, error) {
	pairs := []function.SubdomainFunctionPair{}

	makeComputer := func() (*Computer, error) {
		pwf, err := function.NewPiecewiseFunction(pairs)
		return &Computer{pwf}, err
	}

	if len(changes) == 0 {
		pairs = append(pairs, function.SubdomainFunctionPair{
			Subdomain: util.NewRange(-math.MaxFloat64, math.MaxFloat64),
			Function:  function.NewConstantFunction(startVal),
		})

		return makeComputer()
	}

	offsets := gatherChangeOffsets(changes)

	// from the beginning of time
	pairs = append(pairs, function.SubdomainFunctionPair{
		Subdomain: util.NewRange(-math.MaxFloat64, offsets[0]),
		Function:  function.NewConstantFunction(startVal),
	})

	n := len(offsets)
	lastEndVal := startVal
	lastChangeEnd := offsets[0]

	for i, offset := range offsets {
		change := changes[offset]
		changeEnd := offset + change.Duration

		if change.Transition != TransitionImmediate {
			// Verify that changes don't overlap
			if i < (n - 1) {
				nextOffset := offsets[i+1]

				if nextOffset < changeEnd {
					err := fmt.Errorf("change at offset %f overlaps with change at offset %f", offset, nextOffset)
					return nil, err
				}
			}

			p1 := function.NewPoint(offset, lastEndVal)
			p2 := function.NewPoint(changeEnd, change.EndValue)

			var f function.Function

			switch change.Transition {
			case TransitionLinear:
				f = function.NewLinearFunctionFromPoints(p1, p2)
			case TransitionSigmoid:
				f = function.NewSigmoidFunction(p1, p2)
			default:
				return nil, fmt.Errorf("unknown transition tpe %d", change.Transition)
			}

			pairs = append(pairs, function.SubdomainFunctionPair{
				Subdomain: util.NewRange(offset, changeEnd),
				Function:  f,
			})
		}

		// Keep constant value from the change end until the next offset
		if i < (n - 1) {
			pairs = append(pairs, function.SubdomainFunctionPair{
				Subdomain: util.NewRange(changeEnd, offsets[i+1]),
				Function:  function.NewConstantFunction(change.EndValue),
			})
		}

		lastEndVal = change.EndValue
		lastChangeEnd = changeEnd
	}

	// until the end of time
	pairs = append(pairs, function.SubdomainFunctionPair{
		Subdomain: util.NewRange(lastChangeEnd, math.MaxFloat64),
		Function:  function.NewConstantFunction(lastEndVal),
	})

	return makeComputer()
}

func gatherChangeOffsets(changes map[float64]*Change) []float64 {
	offsets := []float64{}

	for offset, _ := range changes {
		offsets = append(offsets, offset)
	}

	sort.Float64s(offsets)

	return offsets
}
