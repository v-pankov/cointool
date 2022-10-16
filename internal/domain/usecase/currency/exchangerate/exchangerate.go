package exchangerate

import (
	"context"
	"fmt"

	"github.com/vdrpkv/cointool/internal/domain/entity/currency"
)

type UseCaseGetExchangeRate interface {
	DoUseCaseGetExchangeRate(
		ctx context.Context,
		from, to currency.Symbol,
	) (
		currency.ExchangeRate,
		error,
	)
}

type useCaseGetExchangeRate struct {
	exchangeRateClient ExchangeRateClient
	minRateVal         float64
}

func NewUseCaseGetExchangeRate(
	exchangeRateClient ExchangeRateClient,
	minRateVal float64,
) UseCaseGetExchangeRate {
	return useCaseGetExchangeRate{
		exchangeRateClient: exchangeRateClient,
		minRateVal:         minRateVal,
	}
}

func (u useCaseGetExchangeRate) DoUseCaseGetExchangeRate(
	ctx context.Context,
	from, to currency.Symbol,
) (
	currency.ExchangeRate,
	error,
) {
	rate, err := u.exchangeRateClient.GetExchangeRate(ctx, from, to)
	if err != nil {
		return 0, fmt.Errorf("exchange rate client: %w", err)
	}

	if err := rate.Validate(u.minRateVal); err != nil {
		return 0, fmt.Errorf("invalid exchange rate: %w", err)
	}

	return rate, nil
}
