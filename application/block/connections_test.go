package block_test

import (
	"testing"

	"github.com/jamestunnell/go-musicality/application/block"
	"github.com/stretchr/testify/assert"
)

func TestConnections(t *testing.T) {
	a1 := block.NewAddr("A", "1")
	a2 := block.NewAddr("A", "2")
	b1 := block.NewAddr("B", "1")
	b2 := block.NewAddr("B", "2")
	c1 := block.NewAddr("C", "1")
	c2 := block.NewAddr("C", "2")

	conns := block.NewConnections()

	assert.Equal(t, 0, conns.Len())

	addr, found := conns.ConnectedInput(a2)

	assert.Nil(t, addr)
	assert.False(t, found)

	addr, found = conns.ConnectedOutput(a1)

	assert.Nil(t, addr)
	assert.False(t, found)

	assert.True(t, conns.Connect(a2, b1))

	assert.False(t, conns.Connect(a2, c1))

	assert.True(t, conns.Connect(b2, c1))
	assert.True(t, conns.Connect(c2, a1))

	assert.Equal(t, 3, conns.Len())

	verifyConnected(t, conns, a2, b1)
	verifyConnected(t, conns, b2, c1)
	verifyConnected(t, conns, c2, a1)
}

func verifyConnected(t *testing.T, conns *block.Connections, out, in *block.Addr) {
	inActual, ok := conns.ConnectedInput(out)

	assert.True(t, ok)
	assert.Equal(t, in.String(), inActual.String())

	outActual, ok := conns.ConnectedOutput(in)

	assert.True(t, ok)
	assert.Equal(t, out.String(), outActual.String())
}
