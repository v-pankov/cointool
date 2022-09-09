// Package rate implements 'cointool rate' command handler accepting string arguments.
package rate

import (
	"context"

	"github.com/vdrpkv/cointool/internal/currency"
	"github.com/vdrpkv/cointool/internal/handler/generic"

	rateHandler "github.com/vdrpkv/cointool/internal/handler/rate"
)

type genericHandler struct {
	rateHandler rateHandler.RateCommandHandler
}

var _ generic.GenericCommandHandler = genericHandler{}

func New(
	rateHandler rateHandler.RateCommandHandler,
) generic.GenericCommandHandler {
	return genericHandler{
		rateHandler: rateHandler,
	}
}

func (h genericHandler) HandleGenericCommand(
	ctx context.Context,
	args []string,
) (
	interface{},
	error,
) {
	if len(args) != 2 {
		return nil, generic.ErrUnexpectedNumberOfArguments
	}

	rate, err := h.rateHandler.HandleRateCommand(
		ctx,
		currency.Symbol(args[0]),
		currency.Symbol(args[1]),
	)

	if err != nil {
		return 0, err
	}

	return rate, nil
}
