package exchangerate

import (
	"context"

	"github.com/vdrpkv/cointool/internal/controller/cli"
	"github.com/vdrpkv/cointool/internal/domain/entity"
	"github.com/vdrpkv/cointool/internal/domain/usecase/exchangerate"
)

type rateController struct {
	useCase exchangerate.UseCaseGetExchangeRate
}

func New(useCase exchangerate.UseCaseGetExchangeRate) cli.CliController {
	return rateController{useCase: useCase}
}

func (c rateController) ExecCliController(ctx context.Context, args []string) (interface{}, error) {
	if len(args) != 2 {
		return nil, cli.ErrUnexpectedNumberOfArguments
	}

	rate, err := c.useCase.DoUseCaseGetExchangeRate(
		ctx,
		entity.CurrencySymbol(args[0]),
		entity.CurrencySymbol(args[1]),
	)

	if err != nil {
		return nil, err
	}

	return rate, nil
}
