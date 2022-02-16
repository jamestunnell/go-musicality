package rhythm_test

import (
	"testing"

	"github.com/jamestunnell/go-musicality/generation/rhythm"
	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-musicality/notation/rat"
	"github.com/stretchr/testify/assert"
)

func TestSubdivider(t *testing.T) {
	smallest := rat.New(1, 16)
	s := rhythm.NewSubdivider(meter.FourFour(), smallest)

	testSubdivider(t, s)

	smallest = rat.New(1, 8)
	s = rhythm.NewSubdivider(meter.ThreeFour(), smallest)

	testSubdivider(t, s)

	smallest = rat.New(1, 32)
	s = rhythm.NewSubdivider(meter.SixEight(), smallest)

	testSubdivider(t, s)
}

func testSubdivider(t *testing.T, s *rhythm.Subdivider) {
	r := s.Subdivide()

	r.Visit(func(n *rhythm.Node) {
		assert.True(t, n.Dur.GreaterEqual(s.Smallest))
	})

	r.Print()
}
