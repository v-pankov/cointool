package convert

import (
	"context"
	"fmt"
	"strconv"

	"github.com/vdrpkv/cointool/internal/currency"
	"github.com/vdrpkv/cointool/internal/handler/generic"

	convertHandler "github.com/vdrpkv/cointool/internal/handler/convert"
)

type genericHandler struct {
	convertHandler convertHandler.ConvertCommandHandler
}

var _ generic.GenericCommandHandler = genericHandler{}

func New(
	convertHandler convertHandler.ConvertCommandHandler,
) generic.GenericCommandHandler {
	return genericHandler{
		convertHandler: convertHandler,
	}
}

func (h genericHandler) HandleGenericCommand(
	ctx context.Context,
	args []string,
) (
	interface{},
	error,
) {
	if len(args) != 3 {
		return nil, generic.ErrUnexpectedNumberOfArguments
	}

	argAmount, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid amount: %w", err)
	}

	var (
		argFrom = args[1]
		argTo   = args[2]
	)

	amount, err := h.convertHandler.HandleConvertCommand(
		ctx,
		currency.Amount(argAmount),
		currency.Symbol(argFrom),
		currency.Symbol(argTo),
	)

	if err != nil {
		return 0, err
	}

	return amount, nil
}
