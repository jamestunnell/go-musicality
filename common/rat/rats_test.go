package rat_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/common/rat"
)

func TestRatsEmpty(t *testing.T) {
	rats := rat.Rationals{}

	assert.Empty(t, rats)
	assert.Equal(t, 0, rats.Len())
}

func TestRatsNotEmpty(t *testing.T) {
	r1 := rat.New(3, 7)
	r2 := rat.New(7, 8)
	r3 := rat.New(25, 2)

	rats := rat.Rationals{r1, r3, r2}

	assert.NotEmpty(t, rats)
	assert.Equal(t, 3, rats.Len())

	sort.Sort(rats)

	assert.Equal(t, r1, rats[0])
	assert.Equal(t, r2, rats[1])
	assert.Equal(t, r3, rats[2])
}

func TestRatsUnion(t *testing.T) {
	r1 := rat.New(3, 7)
	r2 := rat.New(7, 8)
	r3 := rat.New(7, 8)
	r4 := rat.New(25, 4)

	rats1 := rat.Rationals{r1, r2}
	rats2 := rat.Rationals{r3, r4}
	rats3 := rats1.Union(rats2)

	assert.Len(t, rats3, 3)
}

func TestRatsEqual(t *testing.T) {
	r1 := rat.New(3, 7)
	r2 := rat.New(7, 8)
	r3 := rat.New(7, 8)

	rats1 := rat.Rationals{r1, r2, r3}

	assert.True(t, rats1.Equal(rats1))
	assert.False(t, rats1.Equal(rat.Rationals{}))
	assert.False(t, rat.Rationals{}.Equal(rats1))

	rats2 := rat.Rationals{r2, r3}

	assert.False(t, rats1.Equal(rats2))

	rats2 = rat.Rationals{r2, r3, r1}

	assert.False(t, rats1.Equal(rats2))
}

func TestRatsSum(t *testing.T) {
	assert.True(t, rat.Rationals{}.Sum().Equal(rat.Zero()))

	r1 := rat.New(3, 4)
	r2 := rat.New(1, 4)

	rats1 := rat.Rationals{r1, r2}

	assert.True(t, rats1.Sum().Equal(rat.New(1, 1)))
}
