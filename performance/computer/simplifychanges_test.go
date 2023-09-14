package computer_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/notation/change"
	"github.com/jamestunnell/go-musicality/performance/computer"
)

func TestSimplifyChangesEmpty(t *testing.T) {
	assert.Empty(t, computer.SimplifyChanges(0.0, change.Changes{}))
}

func TestSimplifyChangesAllSame(t *testing.T) {
	startVal := 2.6
	c1 := change.NewImmediate(big.NewRat(1, 1), startVal)
	c2 := change.NewImmediate(big.NewRat(1, 1), startVal)
	c3 := change.NewImmediate(big.NewRat(1, 1), startVal)
	changes := change.Changes{c1, c2, c3}

	simplified := computer.SimplifyChanges(startVal, changes)

	assert.Len(t, simplified, 0)
}

func TestSimplifyChangesNoneSame(t *testing.T) {
	startVal := 2.6
	c1 := change.NewImmediate(big.NewRat(1, 1), startVal+0.1)
	c2 := change.NewImmediate(big.NewRat(2, 1), startVal-0.1)
	c3 := change.NewImmediate(big.NewRat(3, 1), startVal)
	changes := change.Changes{c1, c2, c3}

	simplified := computer.SimplifyChanges(startVal, changes)

	assert.Len(t, simplified, 3)
}

func TestSimplifyChangesSomeSame(t *testing.T) {
	startVal := 2.6
	c1 := change.NewImmediate(big.NewRat(1, 1), startVal)
	c2 := change.NewImmediate(big.NewRat(2, 1), startVal-0.1)
	c3 := change.NewImmediate(big.NewRat(3, 1), startVal)
	changes := change.Changes{c1, c2, c3}

	simplified := computer.SimplifyChanges(startVal, changes)

	require.Len(t, simplified, 2)

	assert.True(t, rat.IsEqual(simplified[0].Offset, c2.Offset))
	assert.True(t, rat.IsEqual(simplified[1].Offset, c3.Offset))
}
