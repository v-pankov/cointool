package convert

import (
	"context"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/vdrpkv/cointool/cmd/cli/cointool/variables"
	"github.com/vdrpkv/cointool/internal/coinmarketcap"
	"github.com/vdrpkv/cointool/internal/currency"
	"github.com/vdrpkv/cointool/internal/handlers"
)

var Command = &cobra.Command{
	Use:   "convert amount symbol-from symbol-to",
	Short: "Convert currency",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		ctx, ctxCancel := context.WithTimeout(
			cmd.Context(), variables.Timeout,
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

	amount, err := handlers.HandleConvertCurrency(
		ctx,
		coinmarketcap.NewFiatCurrencyRecognizer(
			variables.ApiKey, variables.ApiPrefix,
		),
		coinmarketcap.NewExchangeRateGetter(
			variables.ApiKey, variables.ApiPrefix,
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
