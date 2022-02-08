package main

import (
	"fmt"
	"os"

	"github.com/jamestunnell/go-musicality/commands"
	"github.com/jamestunnell/go-musicality/commands/generate"
	"github.com/jamestunnell/go-musicality/commands/render"
	"github.com/jamestunnell/go-musicality/commands/validate"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("no command")
		os.Exit(1)
	}

	if len(os.Args) == 2 {
		fmt.Println("no subcommand")
		os.Exit(1)
	}

	var cmd commands.Command
	var err error

	cmdName := os.Args[1]
	subCmdName := os.Args[2]
	args := os.Args[3:]

	switch cmdName {
	case "render":
		if cmd, err = render.NewSubcommand(subCmdName, args...); err != nil {
			fmt.Printf("%v\n", err)

			os.Exit(1)
		}
	case "generate":
		if cmd, err = generate.NewSubcommand(subCmdName, args...); err != nil {
			fmt.Printf("%v\n", err)

			os.Exit(1)
		}
	case "validate":
		if cmd, err = validate.NewSubcommand(subCmdName, args...); err != nil {
			fmt.Printf("%v\n", err)

			os.Exit(1)
		}
	default:
		fmt.Printf("unknown command '%s'", cmdName)

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
