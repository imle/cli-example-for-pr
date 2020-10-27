package main

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"cli-example/app_b"
	"cli-example/commands"
	"cli-example/module"
)

var appB = cli.App{
	Name: "app-b",
	Commands: []*cli.Command{
		&app_b.CmdSpecificToB,
		&commands.CmdShared,
	},
}

func main() {
	module.CommandSetup(appB.Commands)

	_ = appB.Run([]string{"/path/arg", "cmd-specific-to-b"})
	fmt.Println()
	_ = appB.Run([]string{"/path/arg", "cmd-shared"})
}
