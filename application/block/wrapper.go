package block

type Wrapper struct {
	Block Block
	Name  string
	Rank  float64
}

func NewWrapper(name string, block Block) *Wrapper {
	return &Wrapper{
		Block: block,
		Name:  name,
		Rank:  0.0,
	}
}
