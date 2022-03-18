package block

type Block interface {
	Params() map[string]*Param
	Ports() map[string]*Port
	Initialize() error
	Configure()
	Process()
}
