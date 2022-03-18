package block

//go:generate mockgen -destination=mocks/mockblock.go . Block

type Block interface {
	Params() map[string]*Param
	Ports() map[string]*Port
	Initialize() error
	Configure()
	Process()
}
