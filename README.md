# go-bps

[![pkg.go.dev][pkg.go.dev-badge]][pkg.go.dev]
[![Test][test-badge]][test]
[![reviewdog][reviewdog-badge]][reviewdog]
[![Releases][release-badge]][release]

`go-bps`is a Go package to operate the basis point.
Handling floating point numbers in programming causes rounding errors.
To avoid this, all numerical calculations are done using basis points (integer only) in this package.

## What's Basis Point

> A per ten thousand sign or basis point (often denoted as bp, often pronounced as "bip" or "beep") is (a difference of) one hundredth of a percent or equivalently one ten thousandth. The related concept of a permyriad is literally one part per ten thousand. Figures are commonly quoted in basis points in finance, especially in fixed income markets.

[from Wikipedia](https://en.wikipedia.org/wiki/Basis_point)

One part per million(ppm) is used as the minimum unit for basis points on this package.

```
1 ppm  = 0.01 basis points = 0.0001 %
```

## Example

```go
package main

import (
	"fmt"

	"go.mercari.io/go-bps/bps"
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

func main() {
	// principal amount is 14999
	var principal int64 = 14999
	// interest rate is 8.0%
	rate1 := bps.NewFromPercentage(8)
	// interest rate is 2.645% = 264.5 basis points = 2645 deci basis points
	rate2 := bps.NewFromDeciBasisPoint(2645)
	// interest rate is 0.045 = 4.5%
	rate3 := bps.MustFromString(".045")

	b1 := NewBill(principal, rate1)
	b2 := NewBill(principal, rate2)
	b3 := NewBill(principal, rate3)
	bills := Bills{b1, b2, b3}

	// interest fee: 14999 * 8% = 1199.92
	// Amounts() returns fee amount as integer that's rounded off decimal floating point
	fmt.Println(b1.CalcInterestFee().Amounts())
	// interest fee: 14999 * 2.645 = 396.72355
	fmt.Println(b2.CalcInterestFee().Amounts())
	// interest fee: 14999 * 4.5% = 674.955
	fmt.Println(b3.CalcInterestFee().Amounts())
	// sum interest fees: 1199.92 + 396.72355 + 674.955 = 2271.59855
	// not equal 1199 + 396 + 674 = 2269
	fmt.Println(bills.CalcInterestFee().Amounts())
	// Output:
	// 1198
	// 396
	// 674
	// 2271
}
```

## References

- [Basis point \- Wikipedia](https://en.wikipedia.org/wiki/Basis_point)
- [Parts\-per notation \- Wikipedia](https://en.wikipedia.org/wiki/Parts-per_notation)


## Commiters

- Motonori IWATA([@iwata](https://github.com/iwata))

## Contribution

Please read the CLA below carefully before submitting your contribution.

https://www.mercari.com/cla/

## License

Copyright 2020 Mercari, Inc.

Licensed under the MIT License.

<!-- badge links -->
[test]: https://github.com/mercari/go-bps/actions?query=workflow%3A%22test+and+coverage%22
[reviewdog]: https://github.com/mercari/go-bps/actions?query=workflow%3Areviewdog
[pkg.go.dev]: https://pkg.go.dev/go.mercari.io/go-bps
[release]: https://github.com/mercari/go-bps/releases/latest

[test-badge]: https://github.com/mercari/go-bps/workflows/test%20and%20coverage/badge.svg
[reviewdog-badge]: https://github.com/mercari/go-bps/workflows/reviewdog/badge.svg
[pkg.go.dev-badge]: https://pkg.go.dev/badge/go.mercari.io/go-bps
[release-badge]: https://img.shields.io/github/release/mercari/go-bps.svg?style=flat&logo=github
