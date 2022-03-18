package block_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/application/block"
	block_mocks "github.com/jamestunnell/go-musicality/application/block/mocks"
)

func TestManagerGetAddRemove(t *testing.T) {
	m := block.NewManager()

	b, found := m.GetBlock("a")

	assert.False(t, found)
	assert.Nil(t, b)

	assert.False(t, m.RemoveBlock("a"))

	ctrl := gomock.NewController(t)
	a := block_mocks.NewMockBlock(ctrl)

	assert.True(t, m.AddBlock("a", a))

	b, found = m.GetBlock("a")

	assert.True(t, found)
	assert.Equal(t, a, b)

	assert.False(t, m.AddBlock("a", a))

	assert.True(t, m.RemoveBlock("a"))

	b, found = m.GetBlock("a")

	assert.False(t, found)
	assert.Nil(t, b)

	assert.True(t, m.AddBlock("a", a))

	b, found = m.GetBlock("a")

	assert.True(t, found)
	assert.Equal(t, a, b)
}
