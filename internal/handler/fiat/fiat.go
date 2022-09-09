package fiat

import (
	"context"

	"github.com/vdrpkv/cointool/internal/currency"
	"github.com/vdrpkv/cointool/internal/handler"
)

type FiatCommandHandler interface {
	HandleFiatCommand(
		ctx context.Context,
		symbol currency.Symbol,
	) (
		bool,
		error,
	)
}

type fiatHandler struct {
	fiatCurrencyRecognizer handler.FiatCurrencyRecognizer
}

var _ FiatCommandHandler = (*fiatHandler)(nil)

func New(
	fiatCurrencyRecognizer handler.FiatCurrencyRecognizer,
) FiatCommandHandler {
	return &fiatHandler{
		fiatCurrencyRecognizer: fiatCurrencyRecognizer,
	}
}

func (h *fiatHandler) HandleFiatCommand(
	ctx context.Context,
	symbol currency.Symbol,
) (
	bool,
	error,
) {
	return handler.HandleRecognizeFiatCurrency(ctx, h.fiatCurrencyRecognizer, symbol)
}
