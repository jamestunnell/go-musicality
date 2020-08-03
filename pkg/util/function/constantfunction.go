package function

import (
	"github.com/jamestunnell/go-musicality/pkg/util"
)

type ConstantFunction struct {
	value float64
}

func NewConstantFunction(val float64) *ConstantFunction {
	return &ConstantFunction{value: val}
}

func (f *ConstantFunction) At(x float64) float64 {
	return f.value
}

func (f *ConstantFunction) Domain() util.Range {
	return DomainAllFloat64
}
