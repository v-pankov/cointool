package exchangerate

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vdrpkv/cointool/internal/domain/entity"
	"github.com/vdrpkv/cointool/internal/mocks"
)

func Test_RateCommandHandler(t *testing.T) {
	var (
		zeroExchangeRateValue = 0.00001

		stubCtx  = context.Background()
		stubFrom = entity.CurrencySymbol("STUB_FROM")
		stubTo   = entity.CurrencySymbol("STUB_TO")
		stubErr  = errors.New("stub error")
	)

	type (
		testCaseGiveStubs struct {
			rate entity.ExchangeRate
			err  error
		}

		testCaseGive struct {
			stubs testCaseGiveStubs
			from  entity.CurrencySymbol
			to    entity.CurrencySymbol
		}

		testCaseWant struct {
			clientNoCalled bool
			rate           entity.ExchangeRate
			err            error
		}

		testCase struct {
			name string
			give testCaseGive
			want testCaseWant
		}
	)

	for _, testCase := range []testCase{
		{
			"any error",
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
			"negative exchange rate error",
			testCaseGive{
				stubs: testCaseGiveStubs{
					rate: -1,
				},
			},
			testCaseWant{
				err: entity.ErrRateTooSmall,
			},
		},
		{
			"zero exchange rate error",
			testCaseGive{
				stubs: testCaseGiveStubs{
					rate: entity.ExchangeRate(zeroExchangeRateValue / 10),
				},
			},
			testCaseWant{
				err: entity.ErrRateTooSmall,
			},
		},
		{
			"same symbols",
			testCaseGive{
				from: stubFrom,
				to:   stubFrom,
			},
			testCaseWant{
				clientNoCalled: true,
				rate:           1,
			},
		},
		{
			"success",
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
			from := stubFrom
			if testCase.give.from != "" {
				from = testCase.give.from
			}

			to := stubTo
			if testCase.give.to != "" {
				to = testCase.give.to
			}

			exchangeRateClientMock := mocks.NewExchangeRateClient(t)
			if !testCase.want.clientNoCalled {
				exchangeRateClientMock.
					On(
						"GetExchangeRate", stubCtx, from, to,
					).
					Return(
						testCase.give.stubs.rate, testCase.give.stubs.err,
					)
			}

			exchangeRateUseCase := NewUseCaseGetExchangeRate(
				exchangeRateClientMock,
				zeroExchangeRateValue,
			)
			gotRate, gotErr := exchangeRateUseCase.DoUseCaseGetExchangeRate(
				stubCtx, from, to,
			)

			require.ErrorIs(t, gotErr, testCase.want.err)
			require.Equal(t, testCase.want.rate, gotRate)
		})
	}
}
