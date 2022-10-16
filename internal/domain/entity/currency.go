package entity

type CurrencyAmount float64

func (a CurrencyAmount) Float64() float64 {
	return float64(a)
}

type CurrencySymbol string

func (c CurrencySymbol) String() string {
	return string(c)
}
