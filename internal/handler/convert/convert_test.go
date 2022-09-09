package convert

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
		testCaseGiveArgs struct {
			amount currency.Amount
		}

		testCaseGiveStubs struct {
			rate currency.ExchangeRate
			err  error
		}

		testCaseGive struct {
			args  testCaseGiveArgs
			stubs testCaseGiveStubs
		}

		testCaseWant struct {
			amount currency.Amount
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
			rateCommandHandlerMock := mocks.NewRateCommandHandler(t)
			rateCommandHandlerMock.
				On(
					"HandleRateCommand", stubCtx, stubFrom, stubTo,
				).
				Return(
					testCase.give.stubs.rate, testCase.give.stubs.err,
				)

			convertCommandHandler := New(rateCommandHandlerMock)
			gotAmount, gotErr := convertCommandHandler.HandleConvertCommand(
				stubCtx, testCase.give.args.amount, stubFrom, stubTo,
			)

			require.ErrorIs(t, gotErr, testCase.want.err)
			require.Equal(t, testCase.want.amount, gotAmount)
		})
	}
}
