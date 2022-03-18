package blocks_test

// import (
// 	"reflect"
// 	"testing"

// 	"github.com/jamestunnell/go-musicality/application/blocks"
// 	"github.com/jamestunnell/go-musicality/common/rat"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// )

// func TestSink(t *testing.T) {
// 	testSink(t, true)
// 	testSink(t, 2.5)
// 	testSink(t, "not empty")
// 	testSink(t, 25)
// }

// func testSink(t *testing.T, val interface{}) {
// 	c := blocks.NewSink(reflect.TypeOf(val))
// 	ins := c.Inputs()

// 	assert.Len(t, c.Processed, 0)

// 	require.Len(t, ins, 1)
// 	require.Contains(t, ins, blocks.InputName)

// 	expectedVal := reflect.Zero(reflect.TypeOf(val))
// 	in := ins[blocks.InputName]

// 	assert.Equal(t, expectedVal, in.CurrentValue)

// 	in.CurrentValue = val

// 	c.Process(rat.Zero())

// 	assert.Len(t, c.Processed, 1)
// }
