package value_test

import (
	"testing"

	"github.com/jamestunnell/go-musicality/common/value"
	"github.com/stretchr/testify/assert"
)

func TestMapFlat(t *testing.T) {
	const testKey = "abc"

	m := value.Map{}

	v, found := m.FindValue(testKey)

	assert.False(t, found)
	assert.Nil(t, v)
	assert.False(t, m.ChangeValue(testKey, 1))
	assert.False(t, m.RemoveValue(testKey))

	// add the initial value
	m[testKey] = 2

	v, found = m.FindValue(testKey)

	assert.True(t, found)
	assert.Equal(t, 2, v)

	// change the value
	assert.True(t, m.ChangeValue(testKey, 1))

	v, found = m.FindValue(testKey)

	assert.True(t, found)
	assert.Equal(t, 1, v)

	// remove the value
	m.RemoveValue(testKey)

	v, found = m.FindValue(testKey)

	assert.False(t, m.ChangeValue(testKey, 1))
	assert.False(t, found)
	assert.Nil(t, v)
}

func TestMapChangeValueNestedMap(t *testing.T) {
	const (
		testKey  = "target"
		startVal = "howdy"
		nextVal  = 7
	)

	m := value.Map{
		"submap": value.Map{
			"subsubmap1": 1,
			"subsubmap2": value.Map{},
		},
	}

	v, found := m.FindValue(testKey)

	assert.False(t, found)
	assert.Nil(t, v)

	// add the value
	m["submap"].(value.Map)["subsubmap2"].(value.Map)[testKey] = startVal

	v, found = m.FindValue(testKey)

	assert.True(t, found)
	assert.Equal(t, startVal, v)

	// change the value
	assert.True(t, m.ChangeValue(testKey, nextVal))

	v, found = m.FindValue(testKey)

	assert.True(t, found)
	assert.Equal(t, nextVal, v)

	// remove the value
	assert.True(t, m.RemoveValue(testKey))

	v, found = m.FindValue(testKey)

	assert.False(t, m.ChangeValue(testKey, 1))
	assert.False(t, found)
	assert.Nil(t, v)
}
