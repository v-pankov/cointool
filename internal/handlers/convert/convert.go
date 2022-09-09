package handlers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/vdrpkv/cointool/internal/currency"
	"github.com/vdrpkv/cointool/internal/handlers"
	"github.com/vdrpkv/cointool/internal/handlers/generic"
)

type ConvertCommandHandler interface {
	generic.GenericCommandHandler

	HandleConvertCommand(
		ctx context.Context,
		amount currency.Amount,
		from, to currency.Symbol,
	) (
		currency.Amount,
		error,
	)
}

type convertHandler struct {
	fiatCurrencyRecognizer handlers.FiatCurrencyRecognizer
	exchangeRateGetter     handlers.ExchangeRateGetter
}

var _ ConvertCommandHandler = (*convertHandler)(nil)

func New(
	fiatCurrencyRecognizer handlers.FiatCurrencyRecognizer,
	exchangeRateGetter handlers.ExchangeRateGetter,

) ConvertCommandHandler {
	return &convertHandler{
		fiatCurrencyRecognizer: fiatCurrencyRecognizer,
		exchangeRateGetter:     exchangeRateGetter,
	}
}

func (h *convertHandler) HandleConvertCommand(
	ctx context.Context,
	amount currency.Amount,
	from, to currency.Symbol,
) (
	currency.Amount,
	error,
) {
	return handlers.HandleConvertCurrency(
		ctx,
		h.fiatCurrencyRecognizer,
		h.exchangeRateGetter,
		amount,
		from, to,
	)
}

func (h *convertHandler) HandleGenericCommand(
	ctx context.Context,
	args []string,
) (
	interface{},
	error,
) {
	if len(args) < 3 {
		return nil, generic.ErrNotEnoughArgs
	}

	argAmount, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid amount: %w", err)
	}

	var (
		argFrom = args[1]
		argTo   = args[2]
	)

	amount, err := h.HandleConvertCommand(
		ctx,
		currency.Amount(argAmount),
		currency.Symbol(argFrom),
		currency.Symbol(argTo),
	)

	if err != nil {
		return 0, err
	}

	return amount, nil
}
