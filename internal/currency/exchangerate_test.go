package currency

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/suite"
)

func Test_ExchangeRate(t *testing.T) {
	suite.Run(t, new(exchangeRateSuite))
}

type exchangeRateSuite struct {
	suite.Suite
}

func (s *exchangeRateSuite) Test_Validate() {
	zeroVal := 0.0001
	for _, testCase := range []struct {
		name string
		rate ExchangeRate
		err  error
	}{
		{
			"negative exchange rate",
			-0.1,
			ErrExchangeRateNegative,
		},
		{
			"zero exchange rate",
			0.00001,
			ErrExchangeRateZero,
		},
		{
			"valid exchange rate",
			0.0099,
			nil,
		},
	} {
		s.Run(testCase.name, func() {
			gotErr := testCase.rate.Validate(zeroVal)
			s.Require().ErrorIs(gotErr, testCase.err)
		})
	}
}

func (s *exchangeRateSuite) Test_Flip() {
	for _, testCase := range []struct {
		give ExchangeRate
		want ExchangeRate
	}{
		{
			1, 1,
		},
		{
			2, 0.5,
		},
		{
			0, ExchangeRate(math.Inf(0)),
		},
	} {
		s.Run(
			fmt.Sprintf("give %.1f, want %.1f", testCase.give, testCase.want),
			func() {
				s.Require().Equal(testCase.give.Flip(), testCase.want)
			},
		)
	}
}

func (s *exchangeRateSuite) Test_Convert() {
	type (
		testCaseGive struct {
			rate   ExchangeRate
			amount Amount
		}

		testCaseWant struct {
			amount Amount
		}

		testCase struct {
			give testCaseGive
			want testCaseWant
		}
	)

	for _, testCase := range []testCase{
		{
			testCaseGive{
				rate:   1,
				amount: 1,
			},
			testCaseWant{
				amount: 1,
			},
		},
		{
			testCaseGive{
				rate:   0.5,
				amount: 1,
			},
			testCaseWant{
				amount: 0.5,
			},
		},
		{
			testCaseGive{
				rate:   1,
				amount: 0,
			},
			testCaseWant{
				amount: 0,
			},
		},
		{
			testCaseGive{
				rate:   0,
				amount: 1,
			},
			testCaseWant{
				amount: 0,
			},
		},
	} {
		s.Run(
			fmt.Sprintf(
				"give rate=%.1f amount=%.1f, want amount=%.1f",
				testCase.give.rate, testCase.give.amount, testCase.want.amount,
			),
			func() {
				s.Require().Equal(
					testCase.give.rate.Convert(testCase.give.amount),
					testCase.want.amount,
				)
			},
		)
	}
}
