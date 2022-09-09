package rate

import (
	"context"
	"fmt"

	"github.com/vdrpkv/cointool/internal/client/currency/exchangerate"
	"github.com/vdrpkv/cointool/internal/currency"
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
	exchangeRateClient exchangerate.ExchangeRateGetter
}

var _ RateCommandHandler = (*rateHandler)(nil)

func New(
	exchangeRateClient exchangerate.ExchangeRateGetter,
) RateCommandHandler {
	return &rateHandler{
		exchangeRateClient: exchangeRateClient,
	}
}

func (h *rateHandler) HandleRateCommand(
	ctx context.Context,
	from, to currency.Symbol,
) (
	currency.ExchangeRate,
	error,
) {
	rate, err := h.exchangeRateClient.GetExchangeRate(ctx, from, to)
	if err != nil {
		return 0, fmt.Errorf("get exchange rate: %w", err)
	}

	return rate, nil
}
