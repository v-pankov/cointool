package fiat

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vdrpkv/cointool/internal/domain/entity"
	"github.com/vdrpkv/cointool/mocks"
)

func Test_UseCaseRecognizeFiatCurrency(t *testing.T) {
	var (
		stubCtx    = context.Background()
		stubSymbol = entity.CurrencySymbol("STUB_SYMBOL")
		stubErr    = errors.New("stub error")
	)

	type (
		testCaseGiveStubs struct {
			fiat bool
			err  error
		}

		testCaseGive struct {
			stubs testCaseGiveStubs
		}

		testCaseWant struct {
			clientNoCalled bool
			fiat           bool
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
			"success",
			testCaseGive{
				stubs: testCaseGiveStubs{
					fiat: false,
				},
			},
			testCaseWant{
				fiat: false,
			},
		},
		{
			"success",
			testCaseGive{
				stubs: testCaseGiveStubs{
					fiat: true,
				},
			},
			testCaseWant{
				fiat: true,
			},
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			fiatCurrencyClientMock := mocks.NewFiatCurrencyClient(t)
			if !testCase.want.clientNoCalled {
				fiatCurrencyClientMock.
					On(
						"RecognizeFiatCurrency", stubCtx, stubSymbol,
					).
					Return(
						testCase.give.stubs.fiat, testCase.give.stubs.err,
					)
			}

			fiatCurrencyUseCase := NewUseCaseRecognizeFiatCurrency(
				fiatCurrencyClientMock,
			)
			gotRate, gotErr := fiatCurrencyUseCase.DoUseCaseRecognizeFiatCurrency(
				stubCtx, stubSymbol,
			)

			require.ErrorIs(t, gotErr, testCase.want.err)
			require.Equal(t, testCase.want.fiat, gotRate)
		})
	}
}
