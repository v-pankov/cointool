package handlers

import (
	"context"

	"github.com/vdrpkv/cointool/internal/currency"
	"github.com/vdrpkv/cointool/internal/handlers"
)

type RateCommandHandler interface {
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
