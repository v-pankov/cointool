package recognizer

import (
	"context"

	"github.com/vdrpkv/cointool/internal/currency"
)

// FiatCurrencyRecognizer recognizes fiat currencies
type FiatCurrencyRecognizer interface {
	RecognizeFiatCurrency(
		ctx context.Context,
		from currency.Symbol,
	) (
		bool,
		error,
	)
}
