package score

const (
	DefaultStartDynamic = 1.0
	DefaultStartTempo   = 120.0
)

func OptStartVolume(dtynamic float64) OptFunc {
	return func(s *Score) {
		s.Start.Dynamic = dtynamic
	}
}

func OptStartTempo(tempo float64) OptFunc {
	return func(s *Score) {
		s.Start.Tempo = tempo
	}
}
