package rhythm_test

import (
	"fmt"
	"testing"

	"github.com/jamestunnell/go-musicality/composition/rhythm"
	"github.com/jamestunnell/go-musicality/notation/rat"
	"github.com/stretchr/testify/assert"
)

func TestNewNode(t *testing.T) {
	testNodeSubdividePreservesRest(t, true)
	testNodeSubdividePreservesRest(t, false)
}

func testNodeSubdividePreservesRest(t *testing.T, rest bool) {
	t.Run(fmt.Sprintf("rest %v", rest), func(t *testing.T) {
		elem := rhythm.NewElement(rat.New(1, 1), rest)
		root := rhythm.NewNode(elem)

		root.Subdivide(2)

		root.VisitTerminal(1, func(n *rhythm.Node) {
			assert.Equal(t, rest, n.Element().Rest)
		})
	})
}

func TestNodeSubdivideN(t *testing.T) {
	elem := rhythm.NewElement(rat.New(1, 1), false)
	root := rhythm.NewNode(elem)

	root.SubdivideN(2, 3)

	testNodeVisitTerminal(t, "subdivided 3 times", root, 3, []string{"1/8", "1/8", "1/8", "1/8", "1/8", "1/8", "1/8", "1/8"})
}

func TestNodeVisit(t *testing.T) {
	elem := rhythm.NewElement(rat.New(1, 1), false)
	root := rhythm.NewNode(elem)

	root.SubdivideN(2, 4)

	count := 0
	root.Visit(func(level int, n *rhythm.Node) bool {
		count++

		return true
	})

	assert.Equal(t, 16+8+4+2+1, count)
}

func TestNodeVisitTerminal(t *testing.T) {
	elem := rhythm.NewElement(rat.New(1, 1), false)
	root := rhythm.NewNode(elem)

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
			s = append(s, n.Element().Duration.String())
		})

		assert.Equal(t, expectedDurStrings, s)
	})
}
