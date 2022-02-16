package play

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/commands"
	"github.com/jamestunnell/go-musicality/commands/play/midilive"
)

func NewSubcommand(subCmdName string, cliArgs ...string) (commands.Command, error) {
	switch subCmdName {
	case midilive.Name:
		return midilive.NewFromArgs(cliArgs...)
	}

	return nil, fmt.Errorf("unknown subcommand '%s'", subCmdName)
}
