package block

type Manager struct {
	blocks map[string]*Wrapper
}

func NewManager() *Manager {
	return &Manager{
		blocks: map[string]*Wrapper{},
	}
}

func (m *Manager) AddBlock(name string, block Block) bool {
	if _, found := m.blocks[name]; found {
		return false
	}

	m.blocks[name] = NewWrapper(name, block)

	return true
}

func (m *Manager) RemoveBlock(name string) bool {
	if _, found := m.blocks[name]; !found {
		return false
	}

	delete(m.blocks, name)

	return true
}

func (m *Manager) GetBlock(name string) (Block, bool) {
	wrapper, found := m.blocks[name]
	if !found {
		return nil, false
	}

	return wrapper.Block, true
}
