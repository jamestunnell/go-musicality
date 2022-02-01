package section

func OptStartDynamic(dynamic float64) OptFunc {
	return func(s *Section) {
		s.StartDynamic = dynamic
	}
}

func OptStartTempo(tempo float64) OptFunc {
	return func(s *Section) {
		s.StartTempo = tempo
	}
}
