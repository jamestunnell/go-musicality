package rhythm

import (
	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-musicality/notation/rat"
)

type Subdivider struct {
	Meter    *meter.Meter
	Smallest rat.Rat
}

func NewSubdivider(met *meter.Meter, smallest rat.Rat) *Subdivider {
	return &Subdivider{Meter: met, Smallest: smallest}
}

func (s *Subdivider) Subdivide() *Node {
	root := NewNode(s.Meter.MeasureDuration())
	beatNumer := s.Meter.BeatDuration.Rat.Num().Uint64()

	root.Subdivide(s.Meter.BeatsPerMeasure, func(sub *Node) {
		if beatNumer > 1 {
			sub.Subdivide(beatNumer, s.subdivideByTwo)
		} else {
			s.subdivideByTwo(sub)
		}
	})

	return root
}

func (s *Subdivider) subdivideByTwo(sub *Node) {
	nextSubDir := sub.Dur.Mul(rat.New(1, 2))
	if nextSubDir.GreaterEqual(s.Smallest) {
		sub.Subdivide(2, s.subdivideByTwo)
	}
}
