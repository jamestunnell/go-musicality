package block

import (
	"fmt"
)

type Manager struct {
	wrappers map[string]*Wrapper
	conns    *Connections
}

func NewManager() *Manager {
	return &Manager{
		wrappers: map[string]*Wrapper{},
		conns:    NewConnections(),
	}
}

func (m *Manager) AddBlock(name string, block Block) bool {
	if _, found := m.wrappers[name]; found {
		return false
	}

	m.wrappers[name] = NewWrapper(name, block)

	return true
}

func (m *Manager) RemoveBlock(name string) bool {
	if _, found := m.wrappers[name]; !found {
		return false
	}

	delete(m.wrappers, name)

	return true
}

func (m *Manager) GetBlock(name string) (Block, bool) {
	wrapper, found := m.wrappers[name]
	if !found {
		return nil, false
	}

	return wrapper.Block, true
}

func (m *Manager) Connect(outAddr, inAddr *Addr) error {
	out, err := m.FindPort(outAddr)
	if err != nil {
		return err
	}

	in, err := m.FindPort(inAddr)
	if err != nil {
		return err
	}

	if out.Type != OutputPort {
		return fmt.Errorf("%s is not an output", outAddr.String())
	}

	if in.Type != InputPort && in.Type != ControlPort {
		return fmt.Errorf("%s is not an input or control", inAddr.String())
	}

	m.conns.Connect(outAddr, inAddr)

	return nil
}

func (m *Manager) FindPort(addr *Addr) (*Port, error) {
	wrapper, found := m.wrappers[addr.Block]
	if !found {
		return nil, fmt.Errorf("failed to find block %s", addr.Block)
	}

	port, found := wrapper.Block.Ports()[addr.Port]
	if !found {
		return nil, fmt.Errorf("failed to find port %s", addr.Port)
	}

	return port, nil
}

func (m *Manager) FullyConnected() bool {
	for blockName, w := range m.wrappers {
		for portName, p := range w.Block.Ports() {
			addr := NewAddr(blockName, portName)

			switch p.Type {
			case InputPort, ControlPort:
				if _, found := m.conns.ConnectedOutput(addr); !found {
					return false
				}
			case OutputPort:
				if _, found := m.conns.ConnectedInput(addr); !found {
					return false
				}
			}
		}
	}

	return true
}

func (m *Manager) OutputOnlyBlocks() []string {
	names := []string{}

	for blockName, w := range m.wrappers {
		ports := w.Block.Ports()
		outputOnly := true

		for _, port := range ports {
			if port.Type == InputPort || port.Type == ControlPort {
				outputOnly = false

				break
			}
		}

		if outputOnly {
			names = append(names, blockName)
		}
	}

	return names
}

func (m *Manager) AssignRanks() error {
	if len(m.wrappers) == 0 {
		return nil
	}

	if !m.FullyConnected() {
		return fmt.Errorf("blocks are not fully connected")
	}

	for blockName, rank := range m.conns.RankBlocks() {
		m.wrappers[blockName].Rank = rank
	}

	// outputOnly := m.OutputOnlyBlocks()
	// if len(outputOnly) == 0 {
	// 	return fmt.Errorf("no output-only blocks")
	// }

	// // Every block starts at 1
	// for _, w := range m.wrappers {
	// 	w.Ordinal = 1
	// }

	// current := outputOnly
	// next := []string{}

	// for _, blockName := range current {
	// 	w := m.wrappers[blockName]
	// 	ports := w.Block.Ports()
	// 	outputs := FilterPorts(ports, OutputPort)

	// 	// this block is an end of a chain
	// 	if len(outputs) == 0 {
	// 		continue
	// 	}

	// 	for _, output := range outputs {

	// 	}
	// }

	return nil
}
