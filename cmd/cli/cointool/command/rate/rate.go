package rate

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/vdrpkv/cointool/cmd/cli/cointool/variable"
	"github.com/vdrpkv/cointool/internal/coinmarketcap"
	"github.com/vdrpkv/cointool/internal/currency"
	"github.com/vdrpkv/cointool/internal/handler"
)

var Command = &cobra.Command{
	Use:   "rate symbol-from symbol-to",
	Short: "Get exchange rate",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ctx, ctxCancel := context.WithTimeout(
			cmd.Context(), variable.Timeout,
		)
		defer ctxCancel()

		rate, err := run(ctx, args)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			return
		}

		fmt.Println(rate)
	},
}

func run(ctx context.Context, args []string) (currency.ExchangeRate, error) {
	rate, err := handler.HandleGetExchangeRate(
		ctx,
		coinmarketcap.NewExchangeRateGetter(
			variable.ApiKey, variable.ApiPrefix,
		),
		currency.Symbol(args[0]),
		currency.Symbol(args[1]),
	)

	if err != nil {
		return 0, err
	}

	return rate, nil
}
