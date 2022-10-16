package convert

import (
	"context"
	"fmt"
	"strconv"

	"github.com/vdrpkv/cointool/internal/controller/cli"
	"github.com/vdrpkv/cointool/internal/domain/entity"
	"github.com/vdrpkv/cointool/internal/domain/usecase/convert"
)

type convertController struct {
	useCase convert.UseCaseConvertCurrency
}

func New(useCase convert.UseCaseConvertCurrency) cli.CliController {
	return convertController{useCase: useCase}
}

func (c convertController) ExecCliController(ctx context.Context, args []string) (interface{}, error) {
	if len(args) != 3 {
		return nil, cli.ErrUnexpectedNumberOfArguments
	}

	inAmount, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid amount: %w", err)
	}

	outAmount, err := c.useCase.DoUseCaseConvertCurrency(
		ctx,
		entity.CurrencyAmount(inAmount),
		entity.CurrencySymbol(args[1]),
		entity.CurrencySymbol(args[2]),
	)

	if err != nil {
		return nil, err
	}

	return outAmount, nil
}
