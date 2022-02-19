package temperleypitches

import (
	"flag"
	"fmt"
	"strings"

	"github.com/jamestunnell/go-musicality/composition/pitchgen"
	"github.com/jamestunnell/go-musicality/notation/pitch"
)

type GenerateTemperley struct {
	FileName, OutDir, Key string
	NumPitches            int
	Seed                  uint64
}

const Name = "temperley-pitches"

func NewFromArgs(cliArgs ...string) (*GenerateTemperley, error) {
	flagSet := flag.NewFlagSet("temperley-pitches", flag.ExitOnError)
	key := flagSet.String("key", "C major", "Root pitch class and major/minor (e.g. E major)")
	numPitches := flagSet.Int("n", 20, "num pitches to generate")
	// noteDur := flagSet.String("d", "1/4", "note duration")
	// outDir := flagSet.String("outdir", "", "output directory")
	// fileName := flagSet.String("fname", "score.json", "score file name")
	seed := flagSet.Uint64("seed", 0, "random seed")

	if err := flagSet.Parse(cliArgs); err != nil {
		return nil, fmt.Errorf("failed to parse midi args: %w", err)
	}

	cmd := &GenerateTemperley{
		// OutDir:     *outDir,
		Key:        *key,
		NumPitches: *numPitches,
		// FileName:   *fileName,
		Seed: *seed,
	}

	return cmd, nil
}

func (cmd *GenerateTemperley) Name() string {
	return Name
}

func (cmd *GenerateTemperley) Execute() error {
	keyStrs := strings.Split(cmd.Key, " ")
	if len(keyStrs) != 2 {
		return fmt.Errorf("key '%s' is not in form 'C major'", cmd.Key)
	}

	tonicStr := keyStrs[0]
	triadStr := keyStrs[1]

	rootSemitone, err := pitch.ParseSemitone(tonicStr)
	if err != nil {
		return fmt.Errorf("failed to parse key root '%s' as semitone: %w", tonicStr, err)
	}

	var g pitchgen.PitchGenerator

	switch triadStr {
	case "major":
		g, err = pitchgen.NewMajorTemperleyGenerator(rootSemitone, cmd.Seed)
		if err != nil {
			return fmt.Errorf("failed to make major temperley generator: %w", err)
		}
	case "minor":
		g, err = pitchgen.NewMinorTemperleyGenerator(rootSemitone, cmd.Seed)
		if err != nil {
			return fmt.Errorf("failed to make minor temperley generator: %w", err)
		}
	default:
		return fmt.Errorf("unknown key triad '%s', should be major' or 'minor'", triadStr)
	}

	// noteDur, ok := new(big.Rat).SetString(cmd.NoteDur)
	// if !ok {
	// 	return fmt.Errorf("invalid note duration '%s', must be in the form a/b", cmd.NoteDur)
	// }

	pitches := pitchgen.MakePitches(cmd.NumPitches, g)

	fmt.Println(pitches.Strings())

	// noteDur := rat.New(1, 4)
	// sec := section.New()

	// i := 0
	// for fullMeasuresLeft := n / 4; fullMeasuresLeft > 0; fullMeasuresLeft-- {
	// 	m := measure.New(met)

	// 	m.PartNotes["myPart"] = []*note.Note{
	// 		note.New(noteDur, pitches[i]),
	// 		note.New(noteDur, pitches[i+1]),
	// 		note.New(noteDur, pitches[i+2]),
	// 		note.New(noteDur, pitches[i+3]),
	// 	}

	// 	sec.Measures = append(sec.Measures, m)

	// 	i += 4
	// }

	// s := score.New()

	// s.Sections["mySection"] = sec
	// s.Program = []string{"mySection"}

	// d, err := json.Marshal(s)
	// if err != nil {
	// 	return fmt.Errorf("failed to marshal score JSON: %w", err)
	// }

	// outDir := cmd.OutDir

	// if outDir == "" {
	// 	outDir = "."
	// }

	// fpath := path.Join(outDir, cmd.FileName)

	// err = ioutil.WriteFile(fpath, d, fs.ModePerm)
	// if err != nil {
	// 	return fmt.Errorf("failed to write score file: %w", err)
	// }

	return nil
}
