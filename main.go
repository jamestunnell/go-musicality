package main

import (
	"fmt"
	"os"

	"github.com/jamestunnell/go-musicality/commands"
	"github.com/jamestunnell/go-musicality/commands/generate/temperleypitches"
	"github.com/jamestunnell/go-musicality/commands/midismf"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("no command")
		os.Exit(1)
	}

	var cmd commands.Command
	var err error

	switch os.Args[1] {
	case midismf.Name:
		if cmd, err = midismf.NewFromArgs(os.Args[2:]...); err != nil {
			fmt.Printf("%v\n", err)

			os.Exit(1)
		}
	case "generate-temperley-pitches":
		if cmd, err = temperleypitches.NewFromArgs(os.Args[2:]...); err != nil {
			fmt.Printf("%v\n", err)

			os.Exit(1)
		}
	default:
		fmt.Println("expected command 'midi-smf' or 'generate-temperley-pitches'")

		os.Exit(1)
	}

	if err := cmd.Execute(); err != nil {
		fmt.Printf("%s command failed: %v\n", cmd.Name(), err)

		os.Exit(1)
	}

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
