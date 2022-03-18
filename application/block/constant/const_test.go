package constant_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/application/block"
	"github.com/jamestunnell/go-musicality/application/block/constant"
	"github.com/jamestunnell/go-setting/value"
)

func TestConst(t *testing.T) {
	testConst(t, constant.NewBool(), false, value.NewBool(true))
	testConst(t, constant.NewFloat(), 0.0, value.NewFloat(2.5))
	testConst(t, constant.NewInt(), int64(0), value.NewInt(-25))
	testConst(t, constant.NewUInt(), uint64(0), value.NewUInt(25))
	testConst(t, constant.NewString(), "", value.NewString("hello"))
}

func testConst(t *testing.T, c *constant.Constant, expectedStartVal interface{}, anotherVal value.Single) {
	ports := c.Ports()

	require.Len(t, ports, 1)
	require.Contains(t, ports, block.OutputName)

	out := ports[block.OutputName]

	assert.Equal(t, expectedStartVal, out.CurrentValue)

	params := c.Params()

	require.Len(t, params, 1)
	require.Contains(t, params, block.ValueName)

	val := params[block.ValueName]

	val.Value = anotherVal

	require.NoError(t, c.Initialize())

	assert.Equal(t, anotherVal.Value(), out.CurrentValue)
}
