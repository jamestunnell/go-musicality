package generate

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/commands"
	"github.com/jamestunnell/go-musicality/commands/generate/temperleypitches"
)

func NewSubcommand(subCmdName string, cliArgs ...string) (commands.Command, error) {
	switch subCmdName {
	case temperleypitches.Name:
		return temperleypitches.NewFromArgs(cliArgs...)
	}

	return nil, fmt.Errorf("unknown subcommand '%s'", subCmdName)
}
