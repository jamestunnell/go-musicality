package block_test

import (
	"testing"

	"github.com/jamestunnell/go-musicality/application/block"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	blockA = "A"
	blockB = "B"
	blockC = "C"
	blockD = "D"
)

func TestConnectionsSimpleLoop(t *testing.T) {
	aOut := block.NewAddr(blockA, "out")
	aIn := block.NewAddr(blockA, "in")
	bOut := block.NewAddr(blockB, "out")
	bIn := block.NewAddr(blockB, "in")
	cOut := block.NewAddr(blockC, "out")
	cIn := block.NewAddr(blockC, "in")

	conns := block.NewConnections()

	assert.Equal(t, 0, conns.Len())

	addr, found := conns.ConnectedInput(aOut)

	assert.Nil(t, addr)
	assert.False(t, found)

	addr, found = conns.ConnectedOutput(aIn)

	assert.Nil(t, addr)
	assert.False(t, found)

	require.True(t, conns.Connect(aOut, bIn))

	assert.False(t, conns.Connect(aOut, cIn))

	require.True(t, conns.Connect(bOut, cIn))
	require.True(t, conns.Connect(cOut, aIn))

	require.Equal(t, 3, conns.Len())

	verifyConnected(t, conns, aOut, bIn)
	verifyConnected(t, conns, bOut, cIn)
	verifyConnected(t, conns, cOut, aIn)

	ranks := conns.RankBlocks()

	require.Contains(t, ranks, blockA)
	require.Contains(t, ranks, blockB)
	require.Contains(t, ranks, blockC)

	assert.Equal(t, ranks[blockA], ranks[blockB])
	assert.Equal(t, ranks[blockA], ranks[blockC])
	assert.Equal(t, ranks[blockB], ranks[blockC])

	t.Logf("block ranks: %v", ranks)
}

func TestConnectionsNoLoops(t *testing.T) {
	aOut1 := block.NewAddr(blockA, "Out1")
	aOut2 := block.NewAddr(blockA, "Out2")
	bOut := block.NewAddr(blockB, "Out")
	cIn1 := block.NewAddr(blockC, "In1")
	cIn2 := block.NewAddr(blockC, "In2")
	cOut := block.NewAddr(blockC, "Out")
	dIn1 := block.NewAddr(blockD, "In1")
	dIn2 := block.NewAddr(blockD, "In2")

	conns := block.NewConnections()

	require.True(t, conns.Connect(aOut1, cIn1))
	require.True(t, conns.Connect(aOut2, dIn1))
	require.True(t, conns.Connect(bOut, cIn2))
	require.True(t, conns.Connect(cOut, dIn2))

	require.Equal(t, 4, conns.Len())

	verifyConnected(t, conns, aOut1, cIn1)
	verifyConnected(t, conns, aOut2, dIn1)
	verifyConnected(t, conns, bOut, cIn2)
	verifyConnected(t, conns, cOut, dIn2)

	ranks := conns.RankBlocks()

	require.Contains(t, ranks, blockA)
	require.Contains(t, ranks, blockB)
	require.Contains(t, ranks, blockC)
	require.Contains(t, ranks, blockD)

	assert.Greater(t, ranks[blockD], ranks[blockC])
	assert.Greater(t, ranks[blockC], ranks[blockA])
	assert.GreaterOrEqual(t, ranks[blockA], ranks[blockB])

	t.Logf("block ranks: %v", ranks)
}

func TestConnectionsComplexNetwork(t *testing.T) {
	aIn := block.NewAddr(blockA, "In")
	aOut1 := block.NewAddr(blockA, "Out1")
	aOut2 := block.NewAddr(blockA, "Out2")
	bIn := block.NewAddr(blockB, "In")
	bOut := block.NewAddr(blockB, "Out")
	cIn1 := block.NewAddr(blockC, "In1")
	cIn2 := block.NewAddr(blockC, "In2")
	dIn := block.NewAddr(blockD, "In")
	dOut1 := block.NewAddr(blockD, "Out1")
	dOut2 := block.NewAddr(blockD, "Out2")

	conns := block.NewConnections()

	require.True(t, conns.Connect(aOut1, cIn1))
	require.True(t, conns.Connect(aOut2, bIn))
	require.True(t, conns.Connect(bOut, dIn))
	require.True(t, conns.Connect(dOut1, cIn2))
	require.True(t, conns.Connect(dOut2, aIn))

	require.Equal(t, 5, conns.Len())

	verifyConnected(t, conns, aOut1, cIn1)
	verifyConnected(t, conns, aOut2, bIn)
	verifyConnected(t, conns, bOut, dIn)
	verifyConnected(t, conns, dOut1, cIn2)
	verifyConnected(t, conns, dOut2, aIn)

	ranks := conns.RankBlocks()

	require.Contains(t, ranks, blockA)
	require.Contains(t, ranks, blockB)
	require.Contains(t, ranks, blockC)
	require.Contains(t, ranks, blockD)

	// I don't know exactly how the ranks will play out,
	// because the network is complex, but C should have the highest rank
	assert.Greater(t, ranks[blockC], ranks[blockA])
	assert.Greater(t, ranks[blockC], ranks[blockB])
	assert.Greater(t, ranks[blockC], ranks[blockD])

	t.Logf("block ranks: %v", ranks)
}

func verifyConnected(t *testing.T, conns *block.Connections, out, in *block.Addr) {
	inActual, ok := conns.ConnectedInput(out)

	assert.True(t, ok)
	assert.Equal(t, in.String(), inActual.String())

	outActual, ok := conns.ConnectedOutput(in)

	assert.True(t, ok)
	assert.Equal(t, out.String(), outActual.String())
}
