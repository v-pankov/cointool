package rate

import (
	"github.com/spf13/cobra"

	"github.com/vdrpkv/cointool/cmd/cli/cointool/command"
	"github.com/vdrpkv/cointool/cmd/cli/cointool/variable"

	coinmarketcapExchangeRateClient "github.com/vdrpkv/cointool/internal/coinmarketcap/pkg/client/exchangerate"
	genericRateHandler "github.com/vdrpkv/cointool/internal/handler/generic/rate"
	rateHandler "github.com/vdrpkv/cointool/internal/handler/rate"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "rate symbol-from symbol-to",
		Short: "Get exchange rate",
		Run: func(cmd *cobra.Command, args []string) {
			command.RunGenericCommandHandler(
				cmd, args,
				genericRateHandler.New(
					rateHandler.New(
						coinmarketcapExchangeRateClient.NewExchangeRateGetter(
							variable.ApiKey, variable.ApiPrefix,
						),
						variable.ExchangeRateZeroValue,
					),
				),
			)
		},
	}
}
