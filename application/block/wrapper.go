package block

type Wrapper struct {
	Block   Block
	Name    string
	Ordinal int
}

func NewWrapper(name string, block Block) *Wrapper {
	return &Wrapper{
		Block:   block,
		Name:    name,
		Ordinal: 0,
	}
}
