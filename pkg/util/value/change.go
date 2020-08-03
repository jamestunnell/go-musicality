package value

import (
	"github.com/jamestunnell/go-musicality/pkg/util"
)

type Transition int

type Change struct {
	EndValue   float64
	Duration   float64
	Transition Transition
}

const (
	TransitionImmediate Transition = iota
	TransitionLinear
	TransitionSigmoid
)

func NewImmediateChange(endVal float64) *Change {
	return &Change{
		EndValue:   endVal,
		Duration:   0.0,
		Transition: TransitionImmediate,
	}
}

func NewLinearChange(endVal, dur float64) (*Change, error) {
	if dur <= 0.0 {
		return nil, util.NewNonPositiveFloatError(dur)
	}

	return &Change{EndValue: endVal, Duration: dur, Transition: TransitionLinear}, nil
}

func NewSigmoidChange(endVal, dur float64) (*Change, error) {
	if dur <= 0.0 {
		return nil, util.NewNonPositiveFloatError(dur)
	}

	return &Change{EndValue: endVal, Duration: dur, Transition: TransitionSigmoid}, nil
}
