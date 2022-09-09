package convert

import (
	"github.com/spf13/cobra"

	"github.com/vdrpkv/cointool/cmd/cli/cointool/command"
	"github.com/vdrpkv/cointool/cmd/cli/cointool/variable"

	coinmarketcapExchangeRateClient "github.com/vdrpkv/cointool/internal/coinmarketcap/pkg/client/exchangerate"
	convertHandler "github.com/vdrpkv/cointool/internal/handler/convert"
	genericConvertHandler "github.com/vdrpkv/cointool/internal/handler/generic/convert"
	rateHandler "github.com/vdrpkv/cointool/internal/handler/rate"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "convert amount symbol-from symbol-to",
		Short: "Convert currency",
		Run: func(cmd *cobra.Command, args []string) {
			command.RunGenericCommandHandler(
				cmd, args,
				genericConvertHandler.New(
					convertHandler.New(
						rateHandler.New(
							coinmarketcapExchangeRateClient.NewExchangeRateGetter(
								variable.ApiKey, variable.ApiPrefix,
							),
							variable.ExchangeRateZeroValue,
						),
					),
				),
			)
		},
	}
}
