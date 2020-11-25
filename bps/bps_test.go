package bps_test

import (
	"fmt"
	"math/big"
	"testing"

	"go.mercari.io/go-bps/bps"
)

func TestBPS_String_Default_BaseUnit(t *testing.T) {
	t.Log("The default BaseUnit is DeciBasisPoint")

	tests := []struct {
		name string
		b    *bps.BPS
		want string
	}{
		{
			"1 ppm presents `0` as string",
			bps.NewFromPPM(big.NewInt(1)),
			"0",
		},
		{
			"1 deci basis point presents `1` as string",
			bps.NewFromDeciBasisPoint(1),
			"1",
		},
		{
			"1 basis point presents `10` as string",
			bps.NewFromBasisPoint(1),
			"10",
		},
		{
			"1 percentage presents `1000` as string",
			bps.NewFromPercentage(1),
			"1000",
		},
		{
			"1 amount presents `100000` as string",
			bps.NewFromAmount(1),
			"100000",
		},
		{
			"nil presents `0` as string",
			&bps.BPS{},
			"0",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.b.String(); got != tt.want {
				t.Errorf("BPS.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleString() {
	// backup
	u := bps.BaseUnit
	// 15%
	b := bps.NewFromPercentage(15)

	// The default BaseUnit is DeciBasisPoint, so output as deci basis points
	fmt.Println(b)

	// Update BaseUnit to output as basis points
	bps.BaseUnit = bps.BasisPoint
	fmt.Println(b)

	// Update BaseUnit to output as percentages
	bps.BaseUnit = bps.Percentage
	fmt.Println(b)

	// Update BaseUnit to output as ppms
	bps.BaseUnit = bps.PPM
	fmt.Println(b)

	// teardown
	bps.BaseUnit = u
	// Output:
	// 15000
	// 1500
	// 15
	// 150000
}
