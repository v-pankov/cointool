package fiat

import (
	"github.com/spf13/cobra"

	"github.com/vdrpkv/cointool/cmd/cli/cointool/command"
	"github.com/vdrpkv/cointool/cmd/cli/cointool/variable"
	"github.com/vdrpkv/cointool/internal/coinmarketcap"

	fiatHandler "github.com/vdrpkv/cointool/internal/handler/fiat"
	genericFiatHandler "github.com/vdrpkv/cointool/internal/handler/generic/fiat"
)

var Command = &cobra.Command{
	Use:   "fiat symbol",
	Short: "Check is currency fiat",
	Run: func(cmd *cobra.Command, args []string) {
		command.RunGenericCommandHandler(
			cmd, args,
			genericFiatHandler.New(
				fiatHandler.New(
					coinmarketcap.NewFiatCurrencyRecognizer(
						variable.ApiKey, variable.ApiPrefix,
					),
				),
			),
		)
	},
}
