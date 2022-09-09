package currency

import "errors"

type ExchangeRate float64

func (r ExchangeRate) Float64() float64 {
	return float64(r)
}

var (
	ErrExchangeRateZero     = errors.New("exchange is zero")
	ErrExchangeRateNegative = errors.New("exchange rate is negative")
)

func (r ExchangeRate) Validate(zeroVal float64) error {
	val := r.Float64()

	if val < 0 {
		return ErrExchangeRateNegative
	}

	if val <= zeroVal {
		return ErrExchangeRateZero
	}

	return nil
}

func (r ExchangeRate) Flip() ExchangeRate {
	return 1 / r
}

func (r ExchangeRate) Convert(amount Amount) Amount {
	return Amount(r.Float64() * amount.Float64())
}
