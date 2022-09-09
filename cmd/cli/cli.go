package main

import (
	"fmt"

	"github.com/vdrpkv/cointool/cmd/cli/cointool"
	"github.com/vdrpkv/cointool/cmd/cli/cointool/command/convert"
	"github.com/vdrpkv/cointool/cmd/cli/cointool/command/rate"
)

func main() {
	if err := cointool.SetupConfig("config", "."); err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}

	rootCommand := cointool.NewCommand()
	rootCommand.AddCommand(convert.NewCommand())
	rootCommand.AddCommand(rate.NewCommand())

	rootCommand.Execute()
}
