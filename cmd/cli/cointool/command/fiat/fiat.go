package fiat

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
	Use:   "fiat symbol",
	Short: "Check is currency fiat",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx, ctxCancel := context.WithTimeout(
			cmd.Context(), variable.Timeout,
		)
		defer ctxCancel()

		isFiat, err := run(ctx, args)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			return
		}

		fmt.Println(isFiat)
	},
}

func run(ctx context.Context, args []string) (bool, error) {
	isFiat, err := handler.HandleRecognizeFiatCurrency(
		ctx,
		coinmarketcap.NewFiatCurrencyRecognizer(
			variable.ApiKey, variable.ApiPrefix,
		),
		currency.Symbol(args[0]),
	)

	if err != nil {
		return false, err
	}

	return isFiat, nil
}
