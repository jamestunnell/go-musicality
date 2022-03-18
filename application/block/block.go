package block

import (
	"github.com/jamestunnell/go-musicality/common/rat"
)

type Block interface {
	Params() map[string]*Param
	Ports() map[string]*Port
	Initialize() error
	Configure()
	Process(offset rat.Rat)
}
