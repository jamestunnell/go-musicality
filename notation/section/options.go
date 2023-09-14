package section

import "github.com/jamestunnell/go-musicality/notation/meter"

func OptStartMeter(m *meter.Meter) OptFunc {
	return func(s *Section) {
		s.StartMeter = m
	}
}

func OptStartTempo(tempo float64) OptFunc {
	return func(s *Section) {
		s.StartTempo = tempo
	}
}
