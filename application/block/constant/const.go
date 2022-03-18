package constant

import (
	"github.com/jamestunnell/go-musicality/application/block"
	"github.com/jamestunnell/go-setting/value"
)

type Constant struct {
	Out *block.Port
	Val *block.Param
}

func NewFloat() *Constant {
	return newConstant(value.NewFloat(0.0))
}

func NewBool() *Constant {
	return newConstant(value.NewBool(false))
}

func NewInt() *Constant {
	return newConstant(value.NewInt(0))
}

func NewUInt() *Constant {
	return newConstant(value.NewUInt(0))
}

func NewString() *Constant {
	return newConstant(value.NewString(""))
}

func newConstant(val value.Single) *Constant {
	return &Constant{
		Val: block.NewParam(val),
		Out: block.NewOutput(val.Value()),
	}
}

func (b *Constant) Params() map[string]*block.Param {
	return map[string]*block.Param{
		block.ValueName: b.Val,
	}
}

func (b *Constant) Ports() map[string]*block.Port {
	return map[string]*block.Port{
		block.OutputName: b.Out,
	}
}

func (b *Constant) Initialize() error {
	b.Out.CurrentValue = b.Val.Value.Value()

	return nil
}

func (b *Constant) Configure() {
}

func (b *Constant) Process() {
}
