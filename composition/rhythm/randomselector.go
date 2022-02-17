package rhythm

import (
	"math"

	"github.com/jamestunnell/go-musicality/notation/rat"
	"gonum.org/v1/gonum/stat/distuv"
)

type RandomSelector struct {
	maxLevel distuv.Rander
}

func NewRandomSelector(r distuv.Rander) *RandomSelector {
	return &RandomSelector{maxLevel: r}
}

func (s *RandomSelector) MaxLevelAt(x rat.Rat) int {
	maxLevel := int(math.Round(s.maxLevel.Rand()))

	if maxLevel < 0 {
		return 0
	}

	return maxLevel
}
