package rate

import (
	"github.com/spf13/cobra"

	"github.com/vdrpkv/cointool/cmd/cli/cointool/command"
	"github.com/vdrpkv/cointool/cmd/cli/cointool/variable"
	"github.com/vdrpkv/cointool/internal/coinmarketcap"

	genericRateHandler "github.com/vdrpkv/cointool/internal/handler/generic/rate"
	rateHandler "github.com/vdrpkv/cointool/internal/handler/rate"
)

var Command = &cobra.Command{
	Use:   "rate symbol-from symbol-to",
	Short: "Get exchange rate",
	Run: func(cmd *cobra.Command, args []string) {
		command.RunGenericCommandHandler(
			cmd, args,
			genericRateHandler.New(
				rateHandler.New(
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
