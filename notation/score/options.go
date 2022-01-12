package score

const (
	DefaultStartVolume = 1.0
	DefaultStartTempo  = 120.0
)

func OptStartVolume(vol float64) OptFunc {
	return func(s *Score) {
		s.Start.Volume = vol
	}
}

func OptStartTempo(tempo float64) OptFunc {
	return func(s *Score) {
		s.Start.Tempo = tempo
	}
}
