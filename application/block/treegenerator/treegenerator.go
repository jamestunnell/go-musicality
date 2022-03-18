package treegenerator

import (
	"fmt"
	"math"

	"github.com/rs/zerolog/log"

	"github.com/jamestunnell/go-musicality/application/block"
	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/composition/rhythmgen"
	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-setting/constraint"
	"github.com/jamestunnell/go-setting/value"
)

type TreeGenerator struct {
	SmallestDur              *block.Param
	Meter, MaxLevel, NextDur *block.Port

	smallestDur rat.Rat
	root        *rhythmgen.TreeNode
	visitor     *rhythmgen.TreeVisitor
}

const (
	SmallestDurName = "SmallestDur"
	MaxLevelName    = "MaxLevel"
	MeterName       = "Meter"
	NextDurName     = "NextDur"
)

var (
	fourFour           = meter.FourFour()
	defaultSmallestDur = rat.New(1, 32)
	defaultMaxLevel    = int(2)
)

func New() *TreeGenerator {
	sv := value.NewString(defaultSmallestDur.String())

	return &TreeGenerator{
		SmallestDur: block.NewParam(sv, constraint.NewMinLen(1)),
		Meter:       block.NewControl(fourFour),
		MaxLevel:    block.NewInput(defaultMaxLevel),
		NextDur:     block.NewOutput(rat.Zero()),
		smallestDur: defaultSmallestDur,
	}
}

func (b *TreeGenerator) Params() map[string]*block.Param {
	return map[string]*block.Param{
		SmallestDurName: b.SmallestDur,
	}
}

func (b *TreeGenerator) Ports() map[string]*block.Port {
	return map[string]*block.Port{
		MeterName:    b.Meter,
		MaxLevelName: b.MaxLevel,
		NextDurName:  b.NextDur,
	}
}

func (b *TreeGenerator) Initialize() error {
	s := b.SmallestDur.Value.Value().(string)
	smallestDur, ok := rat.FromString(s)
	if !ok {
		return fmt.Errorf("invalid duration string '%s'", s)
	}

	b.smallestDur = smallestDur

	return nil
}

func (b *TreeGenerator) Configure() {
	met := b.Meter.CurrentValue.(*meter.Meter)

	if result := met.Validate(); result != nil {
		log.Warn().Err(result).Interface("meter", met).Msg("ignoring invalid meter")

		return
	}

	b.root = rhythmgen.SubdivideMeasure(met, b.smallestDur)
	b.visitor = rhythmgen.NewTreeVisitor(b.root)
}

func (b *TreeGenerator) Process() {
	maxLevel := b.getMaxLevel()

	reachedTerminal := false
	onVisit := func(level int, n *rhythmgen.TreeNode) bool {
		if level >= maxLevel || n.Terminal() {
			reachedTerminal = true

			b.NextDur.CurrentValue = n.Duration()

			return false
		}

		return true
	}

	for reachedTerminal {
		b.visitor.VisitNext(onVisit)
	}
}

func (b *TreeGenerator) getMaxLevel() int {
	maxLevelFlt := b.MaxLevel.CurrentValue.(float64)
	maxLevel := int(math.Round(maxLevelFlt))

	if maxLevel < 0 {
		log.Warn().Int("max level", maxLevel).Msg("ignoring negative max level")
	}

	return maxLevel
}
