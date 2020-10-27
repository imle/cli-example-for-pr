package app_a

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"cli-example/module"
)

/* LEGACY START */

type ModuleA struct {
}

func (m *ModuleA) Start() {
	fmt.Println("Start A")
}

func (m *ModuleA) Stop() {
	fmt.Println("Stop A")
}

/* LEGACY STOP */

var CmdSpecificToA = cli.Command{
	Name: "cmd-specific-to-a",
	Metadata: map[string]interface{}{
		module.CLIModuleKey: &ModuleA{},
	},
	Before: func(ctx *cli.Context) error {
		fmt.Println("Before A")
		return nil
	},
	Action: func(ctx *cli.Context) error {
		fmt.Println("Action A")
		return nil
	},
	After: func(ctx *cli.Context) error {
		fmt.Println("After A")
		return nil
	},
}
