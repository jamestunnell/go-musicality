package block

import (
	"github.com/jamestunnell/go-setting/constraint"
	"github.com/jamestunnell/go-setting/value"
)

type Param struct {
	Value       value.Single
	Constraints []constraint.Constraint
}

func NewParam(val value.Single, constraints ...constraint.Constraint) *Param {
	return &Param{
		Value:       val,
		Constraints: constraints,
	}
}
