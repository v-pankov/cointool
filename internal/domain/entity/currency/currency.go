package currency

import "errors"

type Amount float64

func (a Amount) Float64() float64 {
	return float64(a)
}

type Symbol string

func (c Symbol) String() string {
	return string(c)
}

type ExchangeRate float64

func (r ExchangeRate) Float64() float64 {
	return float64(r)
}

func (r ExchangeRate) Validate(minVal float64) error {
	val := r.Float64()

	if val <= minVal {
		return ErrRateTooSmall
	}

	return nil
}

var (
	ErrRateTooSmall = errors.New("exchange is too small")
)

func (r ExchangeRate) Flip() ExchangeRate {
	return 1 / r
}

func (r ExchangeRate) Convert(amount Amount) Amount {
	return Amount(r.Float64() * amount.Float64())
}
