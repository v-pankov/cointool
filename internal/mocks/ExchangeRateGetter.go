// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	currency "github.com/vdrpkv/cointool/internal/currency"

	mock "github.com/stretchr/testify/mock"
)

// ExchangeRateGetter is an autogenerated mock type for the ExchangeRateGetter type
type ExchangeRateGetter struct {
	mock.Mock
}

// GetExchangeRate provides a mock function with given fields: ctx, from, to
func (_m *ExchangeRateGetter) GetExchangeRate(ctx context.Context, from currency.Symbol, to currency.Symbol) (currency.ExchangeRate, error) {
	ret := _m.Called(ctx, from, to)

	var r0 currency.ExchangeRate
	if rf, ok := ret.Get(0).(func(context.Context, currency.Symbol, currency.Symbol) currency.ExchangeRate); ok {
		r0 = rf(ctx, from, to)
	} else {
		r0 = ret.Get(0).(currency.ExchangeRate)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, currency.Symbol, currency.Symbol) error); ok {
		r1 = rf(ctx, from, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewExchangeRateGetter interface {
	mock.TestingT
	Cleanup(func())
}

// NewExchangeRateGetter creates a new instance of ExchangeRateGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewExchangeRateGetter(t mockConstructorTestingTNewExchangeRateGetter) *ExchangeRateGetter {
	mock := &ExchangeRateGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
