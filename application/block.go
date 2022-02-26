package application

type Block struct {
	Inputs   []*Input
	Outputs  []*Output
	Params   []*Param
	Controls []*Control
}
