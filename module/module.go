package module

import (
	"github.com/urfave/cli/v2"
)

const CLIModuleKey = "sq.module"

type Module interface {
	Start()
	Stop()
}

// This method helps us avoid writing the startup code for each and every command we write
// which will always remain the same
func CommandSetup(commands []*cli.Command) {
	for _, cmd := range commands {
		if cmd.Metadata != nil && cmd.Metadata[CLIModuleKey] != nil {
			// Must save references in this context, cmd will point to last cmd of loop otherwise
			module := cmd.Metadata[CLIModuleKey].(Module)
			action := cmd.Action

			cmd.Action = func(ctx *cli.Context) error {
				defer StartModule(module).Stop() // This is many-argument, legacy code in our system

				if action != nil {
					return action(ctx)
				}

				return nil
			}
		}

		if len(cmd.Subcommands) > 0 {
			CommandSetup(cmd.Subcommands)
		}
	}
}

func StartModule(mod Module) Module {
	// top level defer recover

	// Legacy startup code

	mod.Start()

	return mod
}
