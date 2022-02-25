package rhythmgen_test

import (
	"testing"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/composition/rhythmgen"
	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/stretchr/testify/assert"
)

func TestSubdivideMeasure(t *testing.T) {
	met := meter.SixEight()
	smallest := rat.New(1, 16)
	root := rhythmgen.SubdivideMeasure(met, smallest)

	verifyTerminalDurs(t, root, 0, []string{"3/4"})
	verifyTerminalDurs(t, root, 1, []string{"3/8", "3/8"})
	verifyTerminalDurs(t, root, 2, []string{"1/8", "1/8", "1/8", "1/8", "1/8", "1/8"})
	verifyTerminalDurs(t, root, root.Depth(),
		[]string{"1/16", "1/16", "1/16", "1/16", "1/16", "1/16", "1/16", "1/16", "1/16", "1/16", "1/16", "1/16"})
}

func verifyTerminalDurs(t *testing.T, root *rhythmgen.TreeNode, maxLevel int, expected []string) {
	terminalDurs := []string{}
	root.VisitTerminal(maxLevel, func(tn *rhythmgen.TreeNode) {
		terminalDurs = append(terminalDurs, tn.Duration().String())
	})

	assert.Equal(t, expected, terminalDurs)
}
