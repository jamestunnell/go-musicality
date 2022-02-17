package rhythm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/composition/rhythm"
	"github.com/jamestunnell/go-musicality/notation/rat"
)

func TestElementDivideByZero(t *testing.T) {
	e := rhythm.NewElement(rat.New(2, 1), false)

	elems := e.Divide(0)

	assert.Empty(t, elems)
}
