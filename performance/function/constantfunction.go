package function

import "github.com/jamestunnell/go-musicality/notation/rat"

type ConstantFunction struct {
	domain Range
	value  float64
}

func NewConstantFunction(val float64) *ConstantFunction {
	return &ConstantFunction{domain: DomainAll(), value: val}
}

func (f *ConstantFunction) At(x rat.Rat) float64 {
	return f.value
}

func (f *ConstantFunction) Domain() Range {
	return f.domain
}
