package rhythm

import (
	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-musicality/notation/rat"
	"github.com/jamestunnell/go-musicality/performance/function"
)

type Generator struct {
	met     *meter.Meter
	measure *Node
}

type Selector interface {
	MaxLevelAt(x rat.Rat) int
}

func NewGenerator(met *meter.Meter, smallestDur rat.Rat) *Generator {
	measure := NewNode(met.MeasureDuration())

	measure.Subdivide(met.BeatsPerMeasure)

	for _, beat := range measure.Subs() {
		beatNumer := met.BeatDuration.Num().Uint64()
		if beatNumer > 1 {
			beat.Subdivide(beatNumer)
		}
	}

	measure.VisitTerminal(2, func(n *Node) {
		n.SubdivideUntil(2, func(n *Node) bool {
			return n.Duration().Div(rat.FromUint64(2)).GreaterEqual(smallestDur)
		})
	})

	return &Generator{
		met:     met,
		measure: measure,
	}
}

func (g *Generator) Depth() int {
	return g.measure.Depth()
}

func (g *Generator) SmallestDur() rat.Rat {
	return g.measure.SmallestDur()
}

func (g *Generator) MakeMeasure(maxLevelFunction function.Function) rat.Rats {
	return g.Make(g.met.MeasureDuration(), maxLevelFunction)
}

func (g *Generator) Make(dur rat.Rat, maxLevelFunction function.Function) rat.Rats {
	durs := rat.Rats{}
	x := rat.Zero()
	maxLevel := int(maxLevelFunction.At(x))
	done := false

	for durs.Sum().Less(dur) {
		g.measure.Visit(func(level int, n *Node) bool {
			if done {
				return false
			}

			if level >= maxLevel || n.Terminal() {
				durs = append(durs, n.Duration())
				x = x.Add(n.Duration())
				maxLevel = int(maxLevelFunction.At(x))
				done = x.GreaterEqual(dur)

				return false
			}

			return true
		})
	}

	diff := durs.Sum().Sub(dur)
	if diff.Positive() {
		last := durs[len(durs)-1]

		durs[len(durs)-1] = last.Sub(diff)
	}

	return durs
}
