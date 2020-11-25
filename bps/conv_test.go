package bps_test

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"

	"github.com/mercari/go-bps/bps"
)

func TestBPS_Amounts(t *testing.T) {
	tests := []struct {
		name string
		ppm  int64
		want int64
	}{
		{
			"1,000,000 ppms equals 1 amount",
			1000000,
			1,
		},
		{
			"1,999,999 ppms equals 1 amount, round off fractions less than 100,000 ppms",
			1999999,
			1,
		},
		{
			"2,000,000 ppms equals 2 amounts",
			2000000,
			2,
		},
		{
			"999,999 ppms equals zero amounts",
			999999,
			0,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			b := bps.NewFromPPM(big.NewInt(tt.ppm))
			if got := b.Amounts(); got != tt.want {
				t.Errorf("BPS.Amounts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBPS_Percentages(t *testing.T) {
	tests := []struct {
		name string
		ppm  *big.Int
		want *big.Int
	}{
		{
			"1,000,000 ppms equals 100 percentages",
			big.NewInt(1000000),
			big.NewInt(100),
		},
		{
			"1,009,999 ppms equals 100 percentages, round off fractions less than 10,000 ppms",
			big.NewInt(1009999),
			big.NewInt(100),
		},
		{
			"1,010,000 ppms equals 101 percentages",
			big.NewInt(1010000),
			big.NewInt(101),
		},
		{
			"10,000 ppms equals 1 percentage",
			big.NewInt(10000),
			big.NewInt(1),
		},
		{
			"9,999 ppms equals zero percentage",
			big.NewInt(9999),
			big.NewInt(0),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			b := bps.NewFromPPM(tt.ppm)
			if got := b.Percentages(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BPS.Percentages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBPS_BasisPoints(t *testing.T) {
	tests := []struct {
		name string
		ppm  *big.Int
		want *big.Int
	}{
		{
			"1,000,000 ppms equals 10,000 basis points",
			big.NewInt(1000000),
			big.NewInt(10000),
		},
		{
			"1,000,099 ppms equals 10,000 basis points, round off fractions less than 1,000 ppms",
			big.NewInt(1000099),
			big.NewInt(10000),
		},
		{
			"1,001,000 ppms equals 10,001 basis points",
			big.NewInt(1000100),
			big.NewInt(10001),
		},
		{
			"99 ppms equals zero basis points",
			big.NewInt(99),
			big.NewInt(0),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			b := bps.NewFromPPM(tt.ppm)
			if got := b.BasisPoints(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BPS.BasisPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBPS_HalfBasisPoints(t *testing.T) {
	tests := []struct {
		name string
		ppm  *big.Int
		want *big.Int
	}{
		{
			"1,000,000 ppms equals 20,000 half basis points",
			big.NewInt(1000000),
			big.NewInt(20000),
		},
		{
			"1,000,049 ppms equals 20,000 half basis points, round off fractions less than 50 ppms",
			big.NewInt(1000049),
			big.NewInt(20000),
		},
		{
			"1,000,050 ppms equals 20,001 half basis points",
			big.NewInt(1000050),
			big.NewInt(20001),
		},
		{
			"49 ppms equals zero half basis points",
			big.NewInt(49),
			big.NewInt(0),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			b := bps.NewFromPPM(tt.ppm)
			if got := b.HalfBasisPoints(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BPS.HalfBasisPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBPS_DeciBasisPoints(t *testing.T) {
	tests := []struct {
		name string
		ppm  *big.Int
		want *big.Int
	}{
		{
			"1,000,000 ppms equals 100,000 deci basis points",
			big.NewInt(1000000),
			big.NewInt(100000),
		},
		{
			"1,000,009 ppms equals 100,000 deci basis points, round off fractions less than 10 ppms",
			big.NewInt(1000009),
			big.NewInt(100000),
		},
		{
			"1,000,010 ppms equals 100,001 deci basis points",
			big.NewInt(1000010),
			big.NewInt(100001),
		},
		{
			"9 ppms equals zero deci basis points",
			big.NewInt(9),
			big.NewInt(0),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			b := bps.NewFromPPM(tt.ppm)
			if got := b.DeciBasisPoints(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BPS.DeciBasisPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBPS_PPMs(t *testing.T) {
	tests := []struct {
		name string
		ppm  *big.Int
		want *big.Int
	}{
		{
			"1,000 ppms",
			big.NewInt(1000),
			big.NewInt(1000),
		},
		{
			"1 ppms",
			big.NewInt(1),
			big.NewInt(1),
		},
		{
			"5 ppms",
			big.NewInt(5),
			big.NewInt(5),
		},
		{
			"nil equal 0 ppms",
			nil,
			big.NewInt(0),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			b := bps.NewFromPPM(tt.ppm)
			if got := b.PPMs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BPS.PPMs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBPS_Rat(t *testing.T) {
	tests := []struct {
		name string
		b    *bps.BPS
		want *big.Rat
	}{
		{
			"10 ppms = 1 / 100,000",
			bps.NewFromPPM(big.NewInt(10)),
			big.NewRat(1, 100000),
		},
		{
			"8 deci basis points = 8 / 100,000",
			bps.NewFromDeciBasisPoint(8),
			big.NewRat(8, 100000),
		},
		{
			"5 basis points = 5 / 10,000",
			bps.NewFromBasisPoint(5),
			big.NewRat(5, 10000),
		},
		{
			"5 basis points = 1 / 2,000",
			bps.NewFromBasisPoint(5),
			big.NewRat(1, 2000),
		},
		{
			"20 percentages = 1 / 5",
			bps.NewFromPercentage(20),
			big.NewRat(1, 5),
		},
		{
			"3 amounts = 3 / 1",
			bps.NewFromAmount(3),
			big.NewRat(3, 1),
		},
		{
			"nil = 0",
			&bps.BPS{},
			big.NewRat(0, 1),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.b.Rat(); got.Cmp(tt.want) != 0 {
				t.Errorf("BPS.Rat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBPS_Float64(t *testing.T) {
	tests := []struct {
		name      string
		b         *bps.BPS
		wantF     float64
		wantExact bool
	}{
		{
			"1 / 4 can represent as float value exactly",
			bps.NewFromAmount(1).Div(4),
			.25,
			true,
		},
		{
			"1 / 3 cannot represent as float value exactly",
			bps.NewFromAmount(1).Div(3),
			.333333,
			false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotF, gotExact := tt.b.Float64()
			if gotF != tt.wantF {
				t.Errorf("BPS.Float64() gotF = %v, want %v", gotF, tt.wantF)
			}
			if gotExact != tt.wantExact {
				t.Errorf("BPS.Float64() gotExact = %v, want %v", gotExact, tt.wantExact)
			}
		})
	}
}

func ExampleBPS_BaseUnitAmounts() {
	// backup
	u := bps.BaseUnit
	// 15%
	b := bps.NewFromPercentage(15)

	// The default BaseUnit is DeciBasisPoint
	fmt.Println(b.BaseUnitAmounts())

	// BaseUnit is updated by PPM
	bps.BaseUnit = bps.PPM
	fmt.Println(b.BaseUnitAmounts())

	// BaseUnit is updated by HalfBasisPoint
	bps.BaseUnit = bps.HalfBasisPoint
	fmt.Println(b.BaseUnitAmounts())

	// BaseUnit is updated by BasisPoint
	bps.BaseUnit = bps.BasisPoint
	fmt.Println(b.BaseUnitAmounts())

	// BaseUnit is updated by Percentage
	bps.BaseUnit = bps.Percentage
	fmt.Println(b.BaseUnitAmounts())

	// teardown
	bps.BaseUnit = u
	// Output:
	// 15000
	// 150000
	// 3000
	// 1500
	// 15
}
