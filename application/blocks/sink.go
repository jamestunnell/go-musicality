package blocks

// import (
// 	"reflect"

// 	"github.com/jamestunnell/go-musicality/application/block"
// 	"github.com/jamestunnell/go-musicality/common/rat"
// )

// type Sink struct {
// 	id string
// 	in *block.Port

// 	Processed []interface{}
// }

// func NewSink(typ reflect.Type) *Sink {
// 	b := &Sink{id: block.NewID()}

// 	b.in = block.NewPort(reflect.Zero(typ), b.id)

// 	return b
// }

// func (b *Sink) ID() string {
// 	return b.id
// }

// func (b *Sink) Params() map[string]*block.Param {
// 	return map[string]*block.Param{}
// }

// func (b *Sink) Controls() map[string]*block.Port {
// 	return map[string]*block.Port{}
// }

// func (b *Sink) Inputs() map[string]*block.Port {
// 	return map[string]*block.Port{
// 		InputName: b.in,
// 	}
// }

// func (b *Sink) Outputs() map[string]*block.Port {
// 	return map[string]*block.Port{}
// }

// func (b *Sink) Initialize() error {
// 	return nil
// }

// func (b *Sink) Configure() {
// }

// func (b *Sink) Process(offset rat.Rat) {
// 	block.LogIO(b)

// 	b.Processed = append(b.Processed, b.in.CurrentValue)
// }
