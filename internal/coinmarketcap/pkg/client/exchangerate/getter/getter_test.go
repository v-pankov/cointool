package getter

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vdrpkv/cointool/internal/currency"
	"github.com/vdrpkv/cointool/internal/mocks"
)

func Test_Getter(t *testing.T) {
	var (
		stubCtx  = context.Background()
		stubFrom = currency.Symbol("STUB_FROM")
		stubTo   = currency.Symbol("STUB_TO")
		stubErr  = errors.New("stub error")
	)

	type (
		testCaseGiveStubsCurrencyExchangeRateGetter struct {
			rate currency.ExchangeRate
			err  error
		}

		testCaseGiveStubsFiatCurrencyRecognizer struct {
			isFiat bool
			err    error
		}

		testCaseGiveStubs struct {
			currencyExchangeRateGetter testCaseGiveStubsCurrencyExchangeRateGetter
			fiatCurrencyRecognizer     testCaseGiveStubsFiatCurrencyRecognizer
		}

		testCaseGive struct {
			stubs testCaseGiveStubs
		}

		testCaseWantCalls struct {
			currencyExchangeRateGetter bool
			fiatCurrencyRecognizer     bool
		}

		testCaseWant struct {
			calls testCaseWantCalls
			rate  currency.ExchangeRate
			err   error
		}

		testCase struct {
			name string
			give testCaseGive
			want testCaseWant
		}
	)

	for _, testCase := range []testCase{
		{
			"fiat currency recognizer fails",
			testCaseGive{
				stubs: testCaseGiveStubs{
					fiatCurrencyRecognizer: testCaseGiveStubsFiatCurrencyRecognizer{
						err: stubErr,
					},
				},
			},
			testCaseWant{
				calls: testCaseWantCalls{
					fiatCurrencyRecognizer: true,
				},
				err: stubErr,
			},
		},
		{
			"currency exchange getter fails",
			testCaseGive{
				stubs: testCaseGiveStubs{
					currencyExchangeRateGetter: testCaseGiveStubsCurrencyExchangeRateGetter{
						err: stubErr,
					},
				},
			},
			testCaseWant{
				calls: testCaseWantCalls{
					fiatCurrencyRecognizer:     true,
					currencyExchangeRateGetter: true,
				},
				err: stubErr,
			},
		},
		{
			"fiat currency",
			testCaseGive{
				stubs: testCaseGiveStubs{
					currencyExchangeRateGetter: testCaseGiveStubsCurrencyExchangeRateGetter{
						rate: 0.5,
					},
					fiatCurrencyRecognizer: testCaseGiveStubsFiatCurrencyRecognizer{
						isFiat: true,
					},
				},
			},
			testCaseWant{
				calls: testCaseWantCalls{
					fiatCurrencyRecognizer:     true,
					currencyExchangeRateGetter: true,
				},
				rate: 2,
			},
		},
		{
			"non fiat currency",
			testCaseGive{
				stubs: testCaseGiveStubs{
					currencyExchangeRateGetter: testCaseGiveStubsCurrencyExchangeRateGetter{
						rate: 0.5,
					},
				},
			},
			testCaseWant{
				calls: testCaseWantCalls{
					fiatCurrencyRecognizer:     true,
					currencyExchangeRateGetter: true,
				},
				rate: 0.5,
			},
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			fiatCurrencyRecognizerMock := mocks.NewFiatCurrencyRecognizer(t)
			if testCase.want.calls.fiatCurrencyRecognizer {
				fiatCurrencyRecognizerMock.
					On(
						"RecognizeFiatCurrency",
						stubCtx, stubFrom,
					).
					Return(
						testCase.give.stubs.fiatCurrencyRecognizer.isFiat,
						testCase.give.stubs.fiatCurrencyRecognizer.err,
					)
			}

			currencyExchangeRateGetterMock := mocks.NewCurrencyExchangeRateGetter(t)
			if testCase.want.calls.currencyExchangeRateGetter {
				// Flip currency symbols in expected call arguments for fiat currency
				from, to := stubFrom, stubTo
				if testCase.give.stubs.fiatCurrencyRecognizer.isFiat {
					from, to = to, from
				}

				currencyExchangeRateGetterMock.
					On(
						"GetCurrencyExchangeRate",
						stubCtx, from, to,
					).
					Return(
						testCase.give.stubs.currencyExchangeRateGetter.rate,
						testCase.give.stubs.currencyExchangeRateGetter.err,
					)
			}

			getter := getter{
				currencyExchangeRateGetter: currencyExchangeRateGetterMock,
				fiatCurrencyRecognizer:     fiatCurrencyRecognizerMock,
			}

			gotRate, gotErr := getter.GetCurrencyExchangeRate(
				stubCtx, stubFrom, stubTo,
			)

			require.ErrorIs(t, gotErr, testCase.want.err)
			require.Equal(t, testCase.want.rate, gotRate)
		})
	}

}
