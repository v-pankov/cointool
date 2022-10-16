package convert

import (
	"context"

	"github.com/vdrpkv/cointool/internal/domain/entity"
	"github.com/vdrpkv/cointool/internal/domain/usecase/exchangerate"
)

type UseCaseConvertCurrency interface {
	DoUseCaseConvertCurrency(
		ctx context.Context,
		amount entity.CurrencyAmount,
		from, to entity.CurrencySymbol,
	) (
		entity.CurrencyAmount,
		error,
	)
}

type useCaseConvertCurrency struct {
	useCaseGetExchangeRate exchangerate.UseCaseGetExchangeRate
}

func NewUseCaseConvertCurrency(
	useCaseGetExchangeRate exchangerate.UseCaseGetExchangeRate,
) UseCaseConvertCurrency {
	return useCaseConvertCurrency{
		useCaseGetExchangeRate: useCaseGetExchangeRate,
	}
}

func (u useCaseConvertCurrency) DoUseCaseConvertCurrency(
	ctx context.Context,
	amount entity.CurrencyAmount,
	from, to entity.CurrencySymbol,
) (
	entity.CurrencyAmount,
	error,
) {
	rate, err := u.useCaseGetExchangeRate.DoUseCaseGetExchangeRate(ctx, from, to)
	if err != nil {
		return 0, err
	}

	return rate.Convert(amount), nil
}
