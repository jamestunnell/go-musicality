package block

type Wrapper struct {
	Block Block
}

func NewWrapper(b Block) *Wrapper {
	return &Wrapper{
		Block: b,
	}
}
