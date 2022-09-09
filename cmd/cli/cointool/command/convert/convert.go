package convert

import (
	"context"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/vdrpkv/cointool/cmd/cli/cointool/variable"
	"github.com/vdrpkv/cointool/internal/coinmarketcap"
	"github.com/vdrpkv/cointool/internal/currency"
	"github.com/vdrpkv/cointool/internal/handler"
)

var Command = &cobra.Command{
	Use:   "convert amount symbol-from symbol-to",
	Short: "Convert currency",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		ctx, ctxCancel := context.WithTimeout(
			cmd.Context(), variable.Timeout,
		)
		defer ctxCancel()

		amount, err := run(ctx, args)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			return
		}

		fmt.Println(amount)
	},
}

func run(ctx context.Context, args []string) (currency.Amount, error) {
	argAmount, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid amount: %w", err)
	}

	var (
		argFrom = args[1]
		argTo   = args[2]
	)

	amount, err := handler.HandleConvertCurrency(
		ctx,
		coinmarketcap.NewFiatCurrencyRecognizer(
			variable.ApiKey, variable.ApiPrefix,
		),
		coinmarketcap.NewExchangeRateGetter(
			variable.ApiKey, variable.ApiPrefix,
		),
		currency.Amount(argAmount),
		currency.Symbol(argFrom),
		currency.Symbol(argTo),
	)

	if err != nil {
		return 0, err
	}

	return amount, nil
}
