package block_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/application/block"
)

func TestAddr(t *testing.T) {
	a := &block.Addr{}

	assert.False(t, a.Parse("invalid addr"))

	valid := "valid.addr"

	assert.True(t, a.Parse(valid))
	assert.Equal(t, valid, a.String())
}
