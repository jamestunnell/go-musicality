package rhythms

import (
	"math/rand"

	"github.com/jamestunnell/go-musicality/composition/rhythm"
	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-musicality/notation/rat"
)

func RandomMeasure(met *meter.Meter, smallestDur rat.Rat) rhythm.Elements {
	elem := rhythm.NewElement(met.MeasureDuration(), false)
	measure := rhythm.NewNode(elem)

	measure.Subdivide(met.BeatsPerMeasure)

	for _, beat := range measure.Subs() {
		beatNumer := met.BeatDuration.Rat.Num().Uint64()
		if beatNumer > 1 {
			beat.Subdivide(beatNumer)
		}
	}

	measure.VisitTerminal(2, func(n *rhythm.Node) {
		n.SubdivideUntil(2, func(n *rhythm.Node) bool {
			return n.Element().Duration.Div(rat.FromUint64(2)).GreaterEqual(smallestDur)
		})
	})

	elems := rhythm.Elements{}
	depth := measure.Depth()
	maxLevel := 1 + rand.Intn(depth)

	measure.Visit(func(level int, n *rhythm.Node) bool {
		if level == maxLevel || n.Terminal() {
			elems = append(elems, n.Element())

			maxLevel = 1 + rand.Intn(depth)

			return false
		}

		return true
	})

	return elems
}
