package rhythm_test

import (
	"math"
	"testing"

	"github.com/jamestunnell/go-musicality/composition/rhythm"
	"github.com/jamestunnell/go-musicality/notation/rat"
	"github.com/stretchr/testify/assert"
)

func TestNodeSubdivideByZero(t *testing.T) {
	root := rhythm.NewNode(rat.New(1, 1))

	root.Subdivide(0)

	assert.Equal(t, 0, root.Depth())
}

func TestNodeSubdivideN(t *testing.T) {
	root := rhythm.NewNode(rat.New(1, 1))

	root.SubdivideN(2, 3)

	testNodeVisitTerminal(t, "subdivided 3 times", root, 3, []string{"1/8", "1/8", "1/8", "1/8", "1/8", "1/8", "1/8", "1/8"})
}

func TestNodeSubdivideUntil(t *testing.T) {
	root := rhythm.NewNode(rat.New(1, 1))
	smallestDur := rat.New(1, 8)

	root.VisitTerminal(2, func(n *rhythm.Node) {
		n.SubdivideUntil(2, func(n *rhythm.Node) bool {
			subdur := n.Duration().Div(rat.FromUint64(2))
			if subdur.GreaterEqual(smallestDur) {
				return true
			}

			return false
		})
	})

	terminals := []string{}
	root.VisitTerminal(math.MaxInt, func(n *rhythm.Node) {
		terminals = append(terminals, n.Duration().String())
	})

	assert.Equal(t, []string{"1/8", "1/8", "1/8", "1/8", "1/8", "1/8", "1/8", "1/8"}, terminals)
}

func TestNodeVisit(t *testing.T) {
	root := rhythm.NewNode(rat.New(1, 1))

	root.SubdivideN(2, 4)

	count := 0
	root.Visit(func(level int, n *rhythm.Node) bool {
		count++

		return true
	})

	assert.Equal(t, 16+8+4+2+1, count)
}

func TestNodeVisitTerminal(t *testing.T) {
	root := rhythm.NewNode(rat.New(1, 1))

	testNodeVisitTerminal(t, "root only - max level 0", root, 0, []string{"1/1"})
	testNodeVisitTerminal(t, "root only - max level 1", root, 1, []string{"1/1"})

	root.Subdivide(2)

	testNodeVisitTerminal(t, "split once - max level 0", root, 0, []string{"1/1"})
	testNodeVisitTerminal(t, "split once - max level 1", root, 1, []string{"1/2", "1/2"})

	root.Subs()[0].Subdivide(2)
	root.Subs()[1].Subdivide(3)

	testNodeVisitTerminal(t, "split twice - max level 0", root, 0, []string{"1/1"})
	testNodeVisitTerminal(t, "split twice - max level 1", root, 1, []string{"1/2", "1/2"})
	testNodeVisitTerminal(t, "split twice - max level 2", root, 2, []string{"1/4", "1/4", "1/6", "1/6", "1/6"})
}

func testNodeVisitTerminal(t *testing.T, name string, root *rhythm.Node, maxLevel int, expectedDurStrings []string) {
	t.Run(name, func(t *testing.T) {
		s := []string{}

		root.VisitTerminal(maxLevel, func(n *rhythm.Node) {
			s = append(s, n.Duration().String())
		})

		assert.Equal(t, expectedDurStrings, s)
	})
}
