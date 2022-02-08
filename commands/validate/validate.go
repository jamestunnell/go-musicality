package validate

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/commands"
	"github.com/jamestunnell/go-musicality/commands/validate/score"
)

func NewSubcommand(subCmdName string, cliArgs ...string) (commands.Command, error) {
	switch subCmdName {
	case score.Name:
		return score.NewFromArgs(cliArgs...)
	}

	return nil, fmt.Errorf("unknown subcommand '%s'", subCmdName)
}
