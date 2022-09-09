package convert

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
	exchangeRateGetter     handler.ExchangeRateGetter
	fiatCurrencyRecognizer handler.FiatCurrencyRecognizer
}

var _ ConvertCommandHandler = (*convertHandler)(nil)

func New(
	exchangeRateGetter handler.ExchangeRateGetter,
	fiatCurrencyRecognizer handler.FiatCurrencyRecognizer,
) ConvertCommandHandler {
	return &convertHandler{
		exchangeRateGetter:     exchangeRateGetter,
		fiatCurrencyRecognizer: fiatCurrencyRecognizer,
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
