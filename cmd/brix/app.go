package brix

import (
	"fmt"
	"github.com/brix-go/brix/config"
	"github.com/brix-go/brix/internal/command/gen"
	"github.com/brix-go/brix/internal/command/init"
	"github.com/brix-go/brix/internal/command/run"
	"github.com/brix-go/brix/internal/command/upgrade"
	"github.com/spf13/cobra"
)

var text = "A tool for generating golang project"
var coloredText = "\x1B[38;2;66;211;146m" + text + "\x1B[39m"

var CmdRoot = &cobra.Command{
	Use:     "brix",
	Example: "brix new demo-api",
	Short:   fmt.Sprintf("\n                               \n,-----.         ,--.           \n|  |) /_ ,--.--.`--',--.  ,--. \n|  .-.  \\|  .--',--. \\  `'  /  \n|  '--' /|  |   |  | /  /.  \\  \n`------' `--'   `--''--'  '--' \n                               \n%s", coloredText),
	Version: fmt.Sprintf("\n                               \n,-----.         ,--.           \n|  |) /_ ,--.--.`--',--.  ,--. \n|  .-.  \\|  .--',--. \\  `'  /  \n|  '--' /|  |   |  | /  /.  \\  \n`------' `--'   `--''--'  '--' \n                               \nBrix %s - Copyright (c) 2023 Nunu\nReleased under the MIT License.\n\n", config.Version),
}

func init() {
	CmdRoot.AddCommand(new.CmdNew)
	//CmdRoot.AddCommand(create.CmdCreate)
	CmdRoot.AddCommand(run.CmdRun)

	CmdRoot.AddCommand(upgrade.CmdUpgrade)
	gen.CmdCreate.AddCommand(gen.CmdCreateController)
	gen.CmdCreate.AddCommand(gen.CmdCreateService)
	gen.CmdCreate.AddCommand(gen.CmdCreateRepository)
	gen.CmdCreate.AddCommand(gen.CmdCreateDomain)
	gen.CmdCreate.AddCommand(gen.CmdCreateAll)

	//CmdRoot.AddCommand(wire.CmdWire)
	//wire.CmdWire.AddCommand(wire.CmdWireAll)
}

// Execute executes the root command.
func Execute() error {
	return CmdRoot.Execute()
}
