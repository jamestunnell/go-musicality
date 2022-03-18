package blocks_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/application/blocks"
	"github.com/jamestunnell/go-setting/value"
)

func TestConst(t *testing.T) {
	testConst(t, blocks.NewConstBool(), false, value.NewBool(true))
	testConst(t, blocks.NewConstFloat(), 0.0, value.NewFloat(2.5))
	testConst(t, blocks.NewConstInt(), int64(0), value.NewInt(-25))
	testConst(t, blocks.NewConstUInt(), uint64(0), value.NewUInt(25))
	testConst(t, blocks.NewConstString(), "", value.NewString("hello"))
}

func testConst(t *testing.T, c *blocks.Const, expectedStartVal interface{}, anotherVal value.Single) {
	ports := c.Ports()

	require.Len(t, ports, 1)
	require.Contains(t, ports, blocks.OutputName)

	out := ports[blocks.OutputName]

	assert.Equal(t, expectedStartVal, out.CurrentValue)

	params := c.Params()

	require.Len(t, params, 1)
	require.Contains(t, params, blocks.ValueName)

	val := params[blocks.ValueName]

	val.Value = anotherVal

	require.NoError(t, c.Initialize())

	assert.Equal(t, anotherVal.Value(), out.CurrentValue)
}
