package currency

type Amount float64

func (a Amount) Float64() float64 {
	return float64(a)
}

type ExchangeRate float64

func (r ExchangeRate) Float64() float64 {
	return float64(r)
}

func (r ExchangeRate) Flip() ExchangeRate {
	return 1 / r
}

func (r ExchangeRate) Convert(amount Amount) Amount {
	return Amount(r.Float64() * amount.Float64())
}

type Symbol string

func (c Symbol) String() string {
	return string(c)
}
