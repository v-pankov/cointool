package rate

import (
	"context"
	"fmt"

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
	exchangeRateGetter     handler.CryptoCurrencyExchangeRateGetter
	fiatCurrencyRecognizer handler.FiatCurrencyRecognizer
}

var _ RateCommandHandler = (*rateHandler)(nil)

func New(
	exchangeRateGetter handler.CryptoCurrencyExchangeRateGetter,
	fiatCurrencyRecognizer handler.FiatCurrencyRecognizer,
) RateCommandHandler {
	return &rateHandler{
		exchangeRateGetter:     exchangeRateGetter,
		fiatCurrencyRecognizer: fiatCurrencyRecognizer,
	}
}

func (h *rateHandler) HandleRateCommand(
	ctx context.Context,
	from, to currency.Symbol,
) (
	currency.ExchangeRate,
	error,
) {
	// Check is FROM currency fiat one.
	isFiat, err := h.fiatCurrencyRecognizer.RecognizeFiatCurrency(
		ctx, from,
	)

	if err != nil {
		return 0, fmt.Errorf("recognize fiat currency: %w", err)
	}

	// Flip symbols because CryptoCurrencyExchangeRateGetter
	// accepts cryptocurrency symbols only as FROM currency.
	if isFiat {
		from, to = to, from
	}

	rate, err := h.exchangeRateGetter.GetCryptoCurrencyExchangeRate(ctx, from, to)
	if err != nil {
		return 0, fmt.Errorf("get exchange rate: %w", err)
	}

	// Flip exchange rate if first currency is fiat one.
	if isFiat {
		rate = rate.Flip()
	}

	return rate, nil
}
