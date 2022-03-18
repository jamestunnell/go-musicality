package block

type Connections struct {
	outputInput map[string]string
	inputOutput map[string]string
}

func NewConnections() *Connections {
	return &Connections{
		outputInput: map[string]string{},
		inputOutput: map[string]string{},
	}
}

func (conns *Connections) Len() int {
	return len(conns.outputInput)
}

func (conns *Connections) Connect(output, input *PortAddr) bool {
	out := output.String()
	in := input.String()

	if _, found := conns.outputInput[out]; found {
		return false
	}

	conns.outputInput[out] = in
	conns.inputOutput[in] = out

	return true
}

func (conns *Connections) ConnectedInput(output *PortAddr) (*PortAddr, bool) {
	return connected(conns.outputInput, output)
}

func (conns *Connections) ConnectedOutput(input *PortAddr) (*PortAddr, bool) {
	return connected(conns.inputOutput, input)
}

func connected(m map[string]string, addr *PortAddr) (*PortAddr, bool) {
	str, found := m[addr.String()]
	if !found {
		return nil, false
	}

	addr2 := &PortAddr{}

	if !addr2.Parse(str) {
		return nil, false
	}

	return addr2, true
}
