package fiat

import (
	"context"

	"github.com/vdrpkv/cointool/internal/domain/entity"
)

type FiatCurrencyClient interface {
	RecognizeFiatCurrency(ctx context.Context, symbol entity.CurrencySymbol) (bool, error)
}
