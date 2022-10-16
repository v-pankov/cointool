package fiat

import (
	"context"

	"github.com/vdrpkv/cointool/internal/controller/cli"
	"github.com/vdrpkv/cointool/internal/domain/entity"
	"github.com/vdrpkv/cointool/internal/domain/usecase/fiat"
)

type fiatController struct {
	useCase fiat.UseCaseRecognizeFiatCurrency
}

func New(useCase fiat.UseCaseRecognizeFiatCurrency) cli.CliController {
	return fiatController{useCase: useCase}
}

func (c fiatController) ExecCliController(ctx context.Context, args []string) (interface{}, error) {
	if len(args) != 1 {
		return nil, cli.ErrUnexpectedNumberOfArguments
	}

	fiat, err := c.useCase.DoUseCaseRecognizeFiatCurrency(
		ctx,
		entity.CurrencySymbol(args[0]),
	)

	if err != nil {
		return nil, err
	}

	return fiat, nil
}
