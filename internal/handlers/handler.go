package handlers

import (
	"coinconv/internal/currency"
	"context"
	"errors"
	"fmt"
)

func HandleConvertCurrency(
	ctx context.Context,
	fiatCurrencyRecognizer FiatCurrencyRecognizer,
	exchangeRateGetter ExchangeRateGetter,
	amount currency.Amount,
	from, to currency.Symbol,
) (
	currency.Amount,
	error,
) {
	isFiat, err := fiatCurrencyRecognizer.RecognizeFiatCurrency(
		ctx, from,
	)
	if err != nil {
		return currency.Amount(0), fmt.Errorf("recognize fiat currency: %w", err)
	}

	if isFiat {
		from, to = to, from
	}

	rate, err := exchangeRateGetter.GetExchangeRate(ctx, from, to)
	if err != nil {
		return currency.Amount(0), fmt.Errorf("get exchange rate: %w", err)
	}

	if isFiat {
		rate = rate.Flip()
	}

	return rate.Convert(amount), nil
}

func HandleGetExchangeRate(
	ctx context.Context,
	exchangeRateGetter ExchangeRateGetter,
	from, to currency.Symbol,
) (
	currency.ExchangeRate,
	error,
) {
	rate, err := exchangeRateGetter.GetExchangeRate(ctx, from, to)
	if err != nil {
		return currency.ExchangeRate(0), fmt.Errorf("get exchange rate: %w", err)
	}

	return rate, nil
}

func HandleRecognizeFiatCurrency(
	ctx context.Context,
	fiatCurrencyRecognizer FiatCurrencyRecognizer,
	symbol currency.Symbol,
) (
	bool,
	error,
) {
	isFiat, err := fiatCurrencyRecognizer.RecognizeFiatCurrency(ctx, symbol)
	if err != nil {
		return false, fmt.Errorf("recognize fiat currency: %w", err)
	}

	return isFiat, nil
}

var (
	ErrCurrencySymbolNotFound = errors.New("currency symbol not found")
	ErrExchangeRateNotFound   = errors.New("exchange rate not found")
)

type (
	ExchangeRateGetter interface {
		GetExchangeRate(
			ctx context.Context,
			from, to currency.Symbol,
		) (
			currency.ExchangeRate,
			error,
		)
	}

	FiatCurrencyRecognizer interface {
		RecognizeFiatCurrency(
			ctx context.Context,
			symbol currency.Symbol,
		) (
			bool,
			error,
		)
	}
)
