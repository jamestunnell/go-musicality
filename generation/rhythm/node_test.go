package rhythm_test

import (
	"testing"

	"github.com/jamestunnell/go-musicality/generation/rhythm"
	"github.com/jamestunnell/go-musicality/notation/rat"
	"github.com/stretchr/testify/assert"
)

func TestTerminalNodes(t *testing.T) {
	n := rhythm.NewNode(rat.New(1, 1))

	testTerminalNodes(t, n, []string{"1/1"})

	n.Subdivide(5, func(i uint64, sub *rhythm.Node) {
		// subdivide even index nodes
		if (i % 2) == 0 {
			sub.Subdivide(2, rhythm.SubdivideDoNothing)
		}
	})

	testTerminalNodes(t, n, []string{"1/10", "1/10", "1/5", "1/10", "1/10", "1/5", "1/10", "1/10"})

	n.Visit(func(n *rhythm.Node) {})
}

func testTerminalNodes(t *testing.T, root *rhythm.Node, expected []string) {
	actual := []string{}
	for _, n := range root.TerminalNodes() {
		actual = append(actual, n.Str)
	}

	assert.Equal(t, expected, actual)
}
