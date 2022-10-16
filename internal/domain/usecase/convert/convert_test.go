package convert

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vdrpkv/cointool/internal/domain/entity"
	"github.com/vdrpkv/cointool/internal/mocks"
)

func Test_UseCaseConvertCurrency(t *testing.T) {
	var (
		stubCtx  = context.Background()
		stubFrom = entity.CurrencySymbol("STUB_FROM")
		stubTo   = entity.CurrencySymbol("STUB_TO")
		stubErr  = errors.New("stub error")
	)

	type (
		testCaseGiveArgs struct {
			amount entity.CurrencyAmount
		}

		testCaseGiveStubs struct {
			rate entity.ExchangeRate
			err  error
		}

		testCaseGive struct {
			args  testCaseGiveArgs
			stubs testCaseGiveStubs
		}

		testCaseWant struct {
			amount entity.CurrencyAmount
			err    error
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
				args: testCaseGiveArgs{
					amount: 1,
				},
				stubs: testCaseGiveStubs{
					rate: 1,
				},
			},
			testCaseWant{
				amount: 1,
			},
		},
		{
			"amount must be multiplied by two",
			testCaseGive{
				args: testCaseGiveArgs{
					amount: 2,
				},
				stubs: testCaseGiveStubs{
					rate: 2,
				},
			},
			testCaseWant{
				amount: 4,
			},
		},
		{
			"amount must be divided by two",
			testCaseGive{
				args: testCaseGiveArgs{
					amount: 2,
				},
				stubs: testCaseGiveStubs{
					rate: 0.5,
				},
			},
			testCaseWant{
				amount: 1,
			},
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			exchangeRateUseCaseMock := mocks.NewUseCaseGetExchangeRate(t)
			exchangeRateUseCaseMock.
				On(
					"DoUseCaseGetExchangeRate", stubCtx, stubFrom, stubTo,
				).
				Return(
					testCase.give.stubs.rate, testCase.give.stubs.err,
				)

			convertCurrencyUseCase := NewUseCaseConvertCurrency(exchangeRateUseCaseMock)
			gotAmount, gotErr := convertCurrencyUseCase.DoUseCaseConvertCurrency(
				stubCtx, testCase.give.args.amount, stubFrom, stubTo,
			)

			require.ErrorIs(t, gotErr, testCase.want.err)
			require.Equal(t, testCase.want.amount, gotAmount)
		})
	}
}
