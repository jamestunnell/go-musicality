package rhythm

import (
	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-musicality/notation/rat"
)

type Generator struct {
	met     *meter.Meter
	measure *Node
}

type Selector interface {
	MaxLevelAt(x rat.Rat) int
}

func NewGenerator(met *meter.Meter, smallestDur rat.Rat) *Generator {
	elem := NewElement(met.MeasureDuration(), false)
	measure := NewNode(elem)

	measure.Subdivide(met.BeatsPerMeasure)

	for _, beat := range measure.Subs() {
		beatNumer := met.BeatDuration.Rat.Num().Uint64()
		if beatNumer > 1 {
			beat.Subdivide(beatNumer)
		}
	}

	measure.VisitTerminal(2, func(n *Node) {
		n.SubdivideUntil(2, func(n *Node) bool {
			return n.Element().Duration.Div(rat.FromUint64(2)).GreaterEqual(smallestDur)
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

func (g *Generator) MakeMeasure(selector Selector) Elements {
	return g.Make(g.met.MeasureDuration(), selector)
}

func (g *Generator) Make(dur rat.Rat, selector Selector) Elements {
	elems := Elements{}
	x := rat.Zero()
	maxLevel := selector.MaxLevelAt(x)
	done := false

	for elems.Duration().Less(dur) {
		g.measure.Visit(func(level int, n *Node) bool {
			if done {
				return false
			}

			if level >= maxLevel || n.Terminal() {
				elem := &Element{Duration: n.Element().Duration, Rest: false}

				elems = append(elems, elem)
				x = x.Add(elem.Duration)
				maxLevel = selector.MaxLevelAt(x)
				done = x.GreaterEqual(dur)

				return false
			}

			return true
		})
	}

	diff := elems.Duration().Sub(dur)
	if diff.Positive() {
		last := elems[len(elems)-1]

		last.Duration = last.Duration.Sub(diff)
	}

	return elems
}
