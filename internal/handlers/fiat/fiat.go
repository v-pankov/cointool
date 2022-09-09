package handlers

import (
	"context"

	"github.com/vdrpkv/cointool/internal/currency"
	"github.com/vdrpkv/cointool/internal/handlers"
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
	fiatCurrencyRecognizer handlers.FiatCurrencyRecognizer
}

var _ FiatCommandHandler = (*fiatHandler)(nil)

func NewFiatHandler(
	fiatCurrencyRecognizer handlers.FiatCurrencyRecognizer,
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
	return handlers.HandleRecognizeFiatCurrency(ctx, h.fiatCurrencyRecognizer, symbol)
}
