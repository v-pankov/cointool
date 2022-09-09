package rate

import (
	"context"
	"fmt"

	"github.com/vdrpkv/cointool/internal/currency"

	currencyExchangeRateClient "github.com/vdrpkv/cointool/internal/client/currency/exchangerate/getter"
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
	exchangeRateClient    currencyExchangeRateClient.CurrencyExchangeRateGetter
	zeroExchangeRateValue float64
}

var _ RateCommandHandler = (*rateHandler)(nil)

func New(
	exchangeRateClient currencyExchangeRateClient.CurrencyExchangeRateGetter,
	zeroExchangeRateValue float64,
) RateCommandHandler {
	return &rateHandler{
		exchangeRateClient:    exchangeRateClient,
		zeroExchangeRateValue: zeroExchangeRateValue,
	}
}

func (h *rateHandler) HandleRateCommand(
	ctx context.Context,
	from, to currency.Symbol,
) (
	currency.ExchangeRate,
	error,
) {
	// Don't call client when symbols are equal.
	if from == to {
		return 1, nil
	}

	rate, err := h.exchangeRateClient.GetExchangeRate(ctx, from, to)
	if err != nil {
		return 0, fmt.Errorf("get exchange rate: %w", err)
	}

	if err := rate.Validate(h.zeroExchangeRateValue); err != nil {
		return 0, fmt.Errorf("validate exchange rate: %w", err)
	}

	return rate, nil
}
