package handlers

import (
	"context"

	"github.com/vdrpkv/cointool/internal/currency"
	"github.com/vdrpkv/cointool/internal/handlers"
	"github.com/vdrpkv/cointool/internal/handlers/generic"
)

type RateCommandHandler interface {
	generic.GenericCommandHandler

	HandleRateCommand(
		ctx context.Context,
		from, to currency.Symbol,
	) (
		currency.ExchangeRate,
		error,
	)
}

type rateHandler struct {
	exchangeRateGetter handlers.ExchangeRateGetter
}

var _ RateCommandHandler = (*rateHandler)(nil)

func NewRateHandler(
	exchangeRateGetter handlers.ExchangeRateGetter,
) RateCommandHandler {
	return &rateHandler{
		exchangeRateGetter: exchangeRateGetter,
	}
}

func (h *rateHandler) HandleRateCommand(
	ctx context.Context,
	from, to currency.Symbol,
) (
	currency.ExchangeRate,
	error,
) {
	return handlers.HandleGetExchangeRate(ctx, h.exchangeRateGetter, from, to)
}

func (h *rateHandler) HandleGenericCommand(
	ctx context.Context,
	args []string,
) (
	interface{},
	error,
) {
	if len(args) < 2 {
		return nil, generic.ErrNotEnoughArgs
	}

	rate, err := h.HandleRateCommand(
		ctx,
		currency.Symbol(args[0]),
		currency.Symbol(args[1]),
	)

	if err != nil {
		return 0, err
	}

	return rate, nil
}
