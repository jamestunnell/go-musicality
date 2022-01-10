package main

import (
	"fmt"
	"os"

	"github.com/jamestunnell/go-musicality/commands"
	"github.com/jamestunnell/go-musicality/commands/scoretomidi"
)

func main() {
	// s := score.New(score.OptStartVolume(0.75))

	// intro := makeIntro()

	// s.Sections = append(s.Sections, intro)

	// validateAndPrintJSON(s)
	if len(os.Args) == 1 {
		fmt.Println("no command")
		os.Exit(1)
	}

	var cmd commands.Command
	var err error

	switch os.Args[1] {
	case "midi":
		if cmd, err = scoretomidi.NewFromArgs(os.Args[2:]...); err != nil {
			fmt.Printf("%v\n", err)

			os.Exit(1)
		}
	default:
		fmt.Println("expected command 'midi'")

		os.Exit(1)
	}

	fmt.Printf("executing %s command\n", cmd.Name())

	if err := cmd.Execute(); err != nil {
		fmt.Printf("%s command failed: %v\n", cmd.Name(), err)

		os.Exit(1)
	}

	fmt.Printf("%s command completed successfully\n", cmd.Name())

	os.Exit(0)
}

// func validateAndPrintJSON(s *score.Score) {
// 	if result := s.Validate(); result != nil {
// 		fmt.Println("Score is not valid")

// 		for ctx, errs := range result.ContextErrors() {
// 			for _, err := range errs {
// 				fmt.Printf("%s: %v", ctx, err)
// 			}
// 		}

// 		return
// 	}

// 	d, err := json.MarshalIndent(s, "", "  ")
// 	if err != nil {
// 		log.Error().Err(err).Msg("failed to marshal score")

// 		return
// 	}

// 	fmt.Println(string(d))
// }

// func makeIntro() *section.Section {
// 	intro := section.New("intro")

// 	intro.AppendMeasures(8, meter.New(4, 4))

// 	return intro
// }
