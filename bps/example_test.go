package bps_test

import (
	"fmt"

	"github.com/mercari/go-bps/bps"
)

type Bill struct {
	Principal int64
	Rate      *bps.BPS
}

func NewBill(principal int64, rate *bps.BPS) *Bill {
	return &Bill{Principal: principal, Rate: rate}
}

// CalcInterestFee calculates an interest fee for `b`
func (b *Bill) CalcInterestFee() *bps.BPS {
	return b.Rate.Mul(b.Principal)
}

type Bills []*Bill

func (bl Bills) CalcInterestFee() *bps.BPS {
	var fees []*bps.BPS
	for _, b := range bl {
		fees = append(fees, b.CalcInterestFee())
	}
	return bps.Sum(bps.NewFromAmount(0), fees...)
}

func Example() {
	// principal amount is 14999
	var principal int64 = 14999
	// interest rate is 8.0%
	rate1 := bps.NewFromPercentage(8)
	// interest rate is 2.645% = 264.5 basis points = 2645 deci basis points
	rate2 := bps.NewFromDeciBasisPoint(2645)
	// interest rate is 4.5%
	rate3 := bps.MustFromString(".045")

	b1 := NewBill(principal, rate1)
	b2 := NewBill(principal, rate2)
	b3 := NewBill(principal, rate3)
	bills := Bills{b1, b2, b3}

	// interest fee: 14999 * 8% = 1199.92
	// Amounts() returns fee amount as integer that's rounded off decimal floating point
	fmt.Println(b1.CalcInterestFee().FloatString(2))
	// interest fee: 14999 * 2.645 = 396.72355
	fmt.Println(b2.CalcInterestFee().FloatString(2))
	// interest fee: 14999 * 4.5% = 674.955
	fmt.Println(b3.CalcInterestFee().FloatString(2))
	// sum interest fees: 1199.92 + 396.72355 + 674.955 = 2271.59855
	// not equal 1199 + 396 + 674 = 2269
	fmt.Println(bills.CalcInterestFee().FloatString(0))
	// Output:
	// 1199.92
	// 396.72
	// 674.96
	// 2272
}
