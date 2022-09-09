package currency

type Amount float64

func (a Amount) Float64() float64 {
	return float64(a)
}
