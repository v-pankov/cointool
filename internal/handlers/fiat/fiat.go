package handlers

import (
	"context"

	"github.com/vdrpkv/cointool/internal/currency"
	"github.com/vdrpkv/cointool/internal/handlers"
	"github.com/vdrpkv/cointool/internal/handlers/generic"
)

type FiatCommandHandler interface {
	generic.GenericCommandHandler

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

func (h *fiatHandler) HandleGenericCommand(
	ctx context.Context,
	args []string,
) (
	interface{},
	error,
) {
	if len(args) < 1 {
		return nil, generic.ErrNotEnoughArgs
	}

	isFiat, err := h.HandleFiatCommand(
		ctx,
		currency.Symbol(args[0]),
	)

	if err != nil {
		return false, err
	}

	return isFiat, nil
}
