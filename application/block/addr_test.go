package block_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/application/block"
)

func TestAddr(t *testing.T) {
	a5 := block.NewAddr("a", "5")
	a3 := block.NewAddr("a", "3")
	b5 := block.NewAddr("b", "5")

	assert.True(t, a5.Equal(a5))
	assert.False(t, a5.Equal(a3))
	assert.False(t, a5.Equal(b5))
}
