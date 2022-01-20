package pitch

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	errExpectedSemitone           = errors.New("expected semitone letter")
	errExpectedAccidentalOrOctave = errors.New("expected accidental or octave")
	errExpectedOctave             = errors.New("expected octave")

	baseSemitoneMap = map[rune]int{
		'C': 0,
		'D': 2,
		'E': 4,
		'F': 5,
		'G': 7,
		'A': 9,
		'B': 11,
	}
)

func parse(s string) (octave, semitone int, err error) {
	if len(s) == 0 {
		return 0, 0, errExpectedSemitone
	}

	semitone, found := baseSemitoneMap[rune(s[0])]
	if !found {
		return 0, 0, fmt.Errorf("invalid semitone '%s'", s[:1])
	}

	if len(s) == 1 {
		return 0, 0, errExpectedAccidentalOrOctave
	}

	hasAccidental := false

	switch rune(s[1]) {
	case '#':
		semitone++

		hasAccidental = true
	case 'b':
		semitone--

		hasAccidental = true
	}

	var octaveStr string

	if hasAccidental {
		octaveStr = s[2:]
	} else {
		octaveStr = s[1:]
	}

	if len(octaveStr) == 0 {
		return 0, 0, errExpectedOctave
	}

	octave, err = strconv.Atoi(octaveStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid octave '%s': %w", octaveStr, err)
	}

	return octave, semitone, nil
}
