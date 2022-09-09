package convert

import (
	"github.com/spf13/cobra"

	"github.com/vdrpkv/cointool/cmd/cli/cointool/command"
	"github.com/vdrpkv/cointool/cmd/cli/cointool/variable"
	"github.com/vdrpkv/cointool/internal/coinmarketcap"

	convertHandler "github.com/vdrpkv/cointool/internal/handler/convert"
	genericConvertHandler "github.com/vdrpkv/cointool/internal/handler/generic/convert"
)

var Command = &cobra.Command{
	Use:   "convert amount symbol-from symbol-to",
	Short: "Convert currency",
	Run: func(cmd *cobra.Command, args []string) {
		command.RunGenericCommandHandler(
			cmd, args,
			genericConvertHandler.New(
				convertHandler.New(
					coinmarketcap.NewExchangeRateGetter(
						variable.ApiKey, variable.ApiPrefix,
					),
					coinmarketcap.NewFiatCurrencyRecognizer(
						variable.ApiKey, variable.ApiPrefix,
					),
				),
			),
		)
	},
}
