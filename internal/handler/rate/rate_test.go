package rate

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vdrpkv/cointool/internal/currency"
	"github.com/vdrpkv/cointool/internal/mocks"
)

func Test_RateCommandHandler(t *testing.T) {
	var (
		stubCtx  = context.Background()
		stubFrom = currency.Symbol("STUB_FROM")
		stubTo   = currency.Symbol("STUB_TO")
		stubErr  = errors.New("stub error")
	)

	type (
		testCaseGiveStubs struct {
			rate currency.ExchangeRate
			err  error
		}

		testCaseGive struct {
			stubs testCaseGiveStubs
		}

		testCaseWant struct {
			rate currency.ExchangeRate
			err  error
		}

		testCase struct {
			name string
			give testCaseGive
			want testCaseWant
		}
	)

	for _, testCase := range []testCase{
		{
			"error",
			testCaseGive{
				stubs: testCaseGiveStubs{
					err: stubErr,
				},
			},
			testCaseWant{
				err: stubErr,
			},
		},
		{
			"no error",
			testCaseGive{
				stubs: testCaseGiveStubs{
					rate: 0.99,
				},
			},
			testCaseWant{
				rate: 0.99,
			},
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			exchangeRateGetterMock := mocks.NewExchangeRateGetter(t)
			exchangeRateGetterMock.
				On(
					"GetExchangeRate", stubCtx, stubFrom, stubTo,
				).
				Return(
					testCase.give.stubs.rate, testCase.give.stubs.err,
				)

			rateCommandHandler := New(exchangeRateGetterMock)
			gotRate, gotErr := rateCommandHandler.HandleRateCommand(
				stubCtx, stubFrom, stubTo,
			)

			require.ErrorIs(t, gotErr, testCase.want.err)
			require.Equal(t, testCase.want.rate, gotRate)
		})
	}
}
