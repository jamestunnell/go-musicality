package blocks

import (
	"github.com/jamestunnell/go-musicality/application/block"
	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-setting/value"
)

type Const struct {
	Out *block.Port
	Val *block.Param
}

func NewConstFloat() *Const {
	return newConst(value.NewFloat(0.0))
}

func NewConstBool() *Const {
	return newConst(value.NewBool(false))
}

func NewConstInt() *Const {
	return newConst(value.NewInt(0))
}

func NewConstUInt() *Const {
	return newConst(value.NewUInt(0))
}

func NewConstString() *Const {
	return newConst(value.NewString(""))
}

func newConst(val value.Single) *Const {
	return &Const{
		Val: block.NewParam(val),
		Out: block.NewOutput(val.Value()),
	}
}

func (b *Const) Params() map[string]*block.Param {
	return map[string]*block.Param{
		ValueName: b.Val,
	}
}

func (b *Const) Ports() map[string]*block.Port {
	return map[string]*block.Port{
		OutputName: b.Out,
	}
}

func (b *Const) Initialize() error {
	b.Out.CurrentValue = b.Val.Value.Value()

	return nil
}

func (b *Const) Configure() {
}

func (b *Const) Process(offset rat.Rat) {
}
