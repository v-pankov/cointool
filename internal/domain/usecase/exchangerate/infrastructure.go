package exchangerate

import (
	"context"

	"github.com/vdrpkv/cointool/internal/domain/entity"
)

type ExchangeRateClient interface {
	GetExchangeRate(ctx context.Context, from, to entity.CurrencySymbol) (entity.ExchangeRate, error)
}
