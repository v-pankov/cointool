package handlers

import (
	"context"

	"github.com/vdrpkv/cointool/internal/currency"
	"github.com/vdrpkv/cointool/internal/handler"
)

type ConvertCommandHandler interface {
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
	fiatCurrencyRecognizer handler.FiatCurrencyRecognizer
	exchangeRateGetter     handler.ExchangeRateGetter
}

var _ ConvertCommandHandler = (*convertHandler)(nil)

func New(
	fiatCurrencyRecognizer handler.FiatCurrencyRecognizer,
	exchangeRateGetter handler.ExchangeRateGetter,

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
	return handler.HandleConvertCurrency(
		ctx,
		h.fiatCurrencyRecognizer,
		h.exchangeRateGetter,
		amount,
		from, to,
	)
}
