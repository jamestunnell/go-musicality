package rat_test

import (
	"testing"

	"github.com/jamestunnell/go-musicality/notation/rat"
	"github.com/stretchr/testify/assert"
)

func TestAccum(t *testing.T) {
	r := rat.Zero()

	r.Accum(rat.New(1, 1))

	assert.True(t, r.Equal(rat.New(1, 1)))

	r.Accum(rat.New(2, 1))

	assert.True(t, r.Equal(rat.New(3, 1)))
}
