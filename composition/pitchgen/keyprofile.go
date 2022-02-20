package pitchgen

import "github.com/jamestunnell/go-musicality/notation/pitch"

type KeyProfile [pitch.SemitonesPerOctave]float64

func (kp KeyProfile) Transpose(n int) KeyProfile {
	n = n % pitch.SemitonesPerOctave

	kpNew := KeyProfile{}

	for i := 0; i < pitch.SemitonesPerOctave; i++ {
		kpNew[i] = kp[(i+n)%12]
	}

	return kpNew
}
