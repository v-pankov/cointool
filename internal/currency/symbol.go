package currency

type Symbol string

func (c Symbol) String() string {
	return string(c)
}
