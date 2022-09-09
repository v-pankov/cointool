package fiat

import (
	"context"

	"github.com/vdrpkv/cointool/internal/currency"
	"github.com/vdrpkv/cointool/internal/handlers/generic"

	fiatHandler "github.com/vdrpkv/cointool/internal/handlers/fiat"
)

type genericHandler struct {
	fiatHandler fiatHandler.FiatCommandHandler
}

var _ generic.GenericCommandHandler = genericHandler{}

func (h genericHandler) HandleGenericCommand(
	ctx context.Context,
	args []string,
) (
	interface{},
	error,
) {
	if len(args) < 1 {
		return nil, generic.ErrNotEnoughArgs
	}

	isFiat, err := h.fiatHandler.HandleFiatCommand(
		ctx,
		currency.Symbol(args[0]),
	)

	if err != nil {
		return false, err
	}

	return isFiat, nil
}
