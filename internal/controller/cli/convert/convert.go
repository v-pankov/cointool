package convert

import (
	"context"
	"fmt"
	"strconv"

	"github.com/vdrpkv/cointool/internal/controller/cli"
	"github.com/vdrpkv/cointool/internal/domain/entity/currency"
	"github.com/vdrpkv/cointool/internal/domain/usecase/currency/convert"
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
		currency.Amount(inAmount),
		currency.Symbol(args[1]),
		currency.Symbol(args[2]),
	)

	if err != nil {
		return nil, err
	}

	return outAmount, nil
}
