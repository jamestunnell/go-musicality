package application

type Named struct {
	name string
}

func (n *Named) Name() string {
	return n.name
}
