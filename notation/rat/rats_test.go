package rat_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/notation/rat"
)

func TestRatsEmpty(t *testing.T) {
	rats := rat.Rats{}

	assert.Empty(t, rats)
	assert.Equal(t, 0, rats.Len())
}

func TestRatsNotEmpty(t *testing.T) {
	r1 := rat.New(3, 7)
	r2 := rat.New(7, 8)
	r3 := rat.New(25, 2)

	rats := rat.Rats{r1, r3, r2}

	assert.NotEmpty(t, rats)
	assert.Equal(t, 3, rats.Len())

	sort.Sort(rats)

	assert.Equal(t, r1, rats[0])
	assert.Equal(t, r2, rats[1])
	assert.Equal(t, r3, rats[2])
}
