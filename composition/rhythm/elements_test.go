package rhythm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/composition/rhythm"
	"github.com/jamestunnell/go-musicality/notation/rat"
)

func TestDuration(t *testing.T) {
	e1 := &rhythm.Element{Duration: rat.New(1, 4)}
	e2 := &rhythm.Element{Duration: rat.New(1, 8)}
	e3 := &rhythm.Element{Duration: rat.New(1, 16)}
	elems := rhythm.Elements{e1, e2, e3}

	assert.True(t, elems.Duration().Equal(rat.New(7, 16)))
}
