package main

import (
	"github.com/vdrpkv/cointool/cmd/cli/cointool"
	"github.com/vdrpkv/cointool/cmd/cli/cointool/command/convert"
	"github.com/vdrpkv/cointool/cmd/cli/cointool/command/fiat"
	"github.com/vdrpkv/cointool/cmd/cli/cointool/command/rate"
)

func init() {
	cointool.Command.AddCommand(convert.Command)
	cointool.Command.AddCommand(fiat.Command)
	cointool.Command.AddCommand(rate.Command)
}

func main() {
	cointool.Command.Execute()
}
