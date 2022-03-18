package block_test

import (
	"testing"

	"github.com/jamestunnell/go-musicality/application/block"
	"github.com/stretchr/testify/assert"
)

func TestConnections(t *testing.T) {
	a1 := block.NewPortAddr("A", "1")
	a2 := block.NewPortAddr("A", "2")
	b1 := block.NewPortAddr("B", "1")
	b2 := block.NewPortAddr("B", "2")
	c1 := block.NewPortAddr("C", "1")
	c2 := block.NewPortAddr("C", "2")

	conns := block.NewConnections()

	assert.Equal(t, 0, conns.Len())

	assert.True(t, conns.Connect(a2, b1))
	assert.True(t, conns.Connect(b2, c1))
	assert.True(t, conns.Connect(c2, a1))

	assert.Equal(t, 3, conns.Len())

	verifyConnected(t, conns, a2, b1)
	verifyConnected(t, conns, b2, c1)
	verifyConnected(t, conns, c2, a1)
}

func verifyConnected(t *testing.T, conns *block.Connections, out, in *block.PortAddr) {
	inActual, ok := conns.ConnectedInput(out)

	assert.True(t, ok)
	assert.Equal(t, in.String(), inActual.String())

	outActual, ok := conns.ConnectedOutput(in)

	assert.True(t, ok)
	assert.Equal(t, out.String(), outActual.String())
}
