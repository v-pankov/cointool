package fiat

import (
	"context"

	"github.com/vdrpkv/cointool/internal/currency"
	"github.com/vdrpkv/cointool/internal/handler"
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
	exchangeRateGetter handler.ExchangeRateGetter
}

var _ RateCommandHandler = (*rateHandler)(nil)

func New(
	exchangeRateGetter handler.ExchangeRateGetter,
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
	return handler.HandleGetExchangeRate(ctx, h.exchangeRateGetter, from, to)
}
