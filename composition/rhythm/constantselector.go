package rhythm

import "github.com/jamestunnell/go-musicality/notation/rat"

type ConstantSelector struct {
	maxLevel int
}

func NewConstantSelector(maxLevel int) *ConstantSelector {
	return &ConstantSelector{maxLevel: maxLevel}
}

func (s *ConstantSelector) MaxLevelAt(x rat.Rat) int {
	return s.maxLevel
}
