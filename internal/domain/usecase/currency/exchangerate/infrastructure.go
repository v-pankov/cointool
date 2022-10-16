package exchangerate

import (
	"context"

	"github.com/vdrpkv/cointool/internal/domain/entity/currency"
)

type ExchangeRateClient interface {
	GetExchangeRate(ctx context.Context, from, to currency.Symbol) (currency.ExchangeRate, error)
}
