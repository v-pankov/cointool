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
	return handler.HandleGetExchangeRate(
		ctx,
		h.fiatCurrencyRecognizer,
		h.exchangeRateGetter,
		from,
		to,
	)
}
