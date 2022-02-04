package sequence

import "github.com/jamestunnell/go-musicality/notation/pitch"

func StepPitches(startPitch, endPitch *pitch.Pitch) pitch.Pitches {
	start := startPitch.TotalSemitone()
	end := endPitch.TotalSemitone()

	if end == start {
		return pitch.Pitches{pitch.New(0, start)}
	}

	var semitones []int

	if end > start {
		semitones = IntRangeAsc(start, end)
	} else {
		semitones = IntRangeDesc(start, end)
	}

	pitches := make(pitch.Pitches, len(semitones))

	for i, semitone := range semitones {
		pitches[i] = pitch.New(0, semitone)
	}

	return pitches
}

func IntRangeAsc(start, end int) []int {
	if start > end {
		return []int{}
	}

	n := (end - start) + 1
	ints := make([]int, n)
	intVal := start

	for i := 0; i < n; i++ {
		ints[i] = intVal

		intVal++
	}

	return ints
}

func IntRangeDesc(start, end int) []int {
	if end > start {
		return []int{}
	}

	n := (start - end) + 1
	ints := make([]int, n)
	intVal := start

	for i := 0; i < n; i++ {
		ints[i] = intVal

		intVal--
	}

	return ints
}
