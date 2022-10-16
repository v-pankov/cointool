package fiat

import (
	"context"

	"github.com/vdrpkv/cointool/internal/domain/entity/currency"
)

type FiatCurrencyClient interface {
	RecognizeFiatCurrency(ctx context.Context, symbol currency.Symbol) (bool, error)
}
