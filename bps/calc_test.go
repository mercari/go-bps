package bps_test

import (
	"math/big"
	"reflect"
	"testing"

	"go.mercari.io/go-bps/bps"
)

func TestBPS_Add(t *testing.T) {
	tests := []struct {
		name string
		b    *bps.BPS
		b2   *bps.BPS
		want *bps.BPS
	}{
		{
			"1 basis point + 1 percentage = 10,100 ppms",
			bps.NewFromBasisPoint(1),
			bps.NewFromPercentage(1),
			bps.NewFromPPM(big.NewInt(10100)),
		},
		{
			"50 ppms + 1 deci basis point = 60 ppms",
			bps.NewFromPPM(big.NewInt(50)),
			bps.NewFromDeciBasisPoint(1),
			bps.NewFromPPM(big.NewInt(60)),
		},
		{
			"1 deci basis point + (-1) basis point = -90 ppms",
			bps.NewFromDeciBasisPoint(1),
			bps.NewFromBasisPoint(-1),
			bps.NewFromPPM(big.NewInt(-90)),
		},
		{
			"nil + 1 ppm = 1 ppms",
			&bps.BPS{},
			bps.NewFromPPM(big.NewInt(1)),
			bps.NewFromPPM(big.NewInt(1)),
		},
		{
			"1 ppm + nil = 1 ppms",
			bps.NewFromPPM(big.NewInt(1)),
			&bps.BPS{},
			bps.NewFromPPM(big.NewInt(1)),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.b.Add(tt.b2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BPS.Add() = %v, want %v", got, tt.want)
			}
			assertImmutableOperation(t, "BPS.Add()", got, tt.b, tt.b2)
		})
	}
}

func assertImmutableOperation(t *testing.T, msg string, got, receiver *bps.BPS, args ...*bps.BPS) {
	t.Helper()

	if &got == &receiver {
		t.Errorf("%s has never mutated the receiver: got=%v, receiver=%v", msg, got, receiver)
	}
	for i, arg := range args {
		if &got == &receiver {
			t.Errorf("%s has never mutated any arguments: got=%v, i=%d, arg=%v", msg, got, i, arg)
		}
	}
}

func TestBPS_Sub(t *testing.T) {
	tests := []struct {
		name string
		b    *bps.BPS
		b2   *bps.BPS
		want *bps.BPS
	}{
		{
			"1 amount - 10 percentages = 900,000 ppms",
			bps.NewFromAmount(1),
			bps.NewFromPercentage(10),
			bps.NewFromPPM(big.NewInt(900000)),
		},
		{
			"1 amount - (-10) percentages = 1100,000 ppms",
			bps.NewFromAmount(1),
			bps.NewFromPercentage(-10),
			bps.NewFromPPM(big.NewInt(1100000)),
		},
		{
			"1 basis point - 10 deci basis point = 0",
			bps.NewFromBasisPoint(1),
			bps.NewFromDeciBasisPoint(10),
			bps.NewFromAmount(0),
		},
		{
			"nil - 1 ppm = -1 ppm",
			&bps.BPS{},
			bps.NewFromPPM(big.NewInt(1)),
			bps.NewFromPPM(big.NewInt(-1)),
		},
		{
			"1 ppm - nil = 1 ppm",
			bps.NewFromPPM(big.NewInt(1)),
			&bps.BPS{},
			bps.NewFromPPM(big.NewInt(1)),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.b.Sub(tt.b2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BPS.Sub() = %v, want %v", got, tt.want)
			}
			assertImmutableOperation(t, "BPS.Sub()", got, tt.b, tt.b2)
		})
	}
}

func TestBPS_Mul(t *testing.T) {
	tests := []struct {
		name string
		b    *bps.BPS
		arg  int64
		want *bps.BPS
	}{
		{
			"1 basis point * 5 = 500 ppms",
			bps.NewFromBasisPoint(1),
			5,
			bps.NewFromPPM(big.NewInt(500)),
		},
		{
			"1 deci basis point * (-10) = -100 ppms",
			bps.NewFromDeciBasisPoint(1),
			-10,
			bps.NewFromPPM(big.NewInt(-100)),
		},
		{
			"-1 percentage * 2 = -20,000 ppms",
			bps.NewFromPercentage(-1),
			2,
			bps.NewFromPPM(big.NewInt(-20000)),
		},
		{
			"nil * 1 = 0",
			&bps.BPS{},
			1,
			bps.NewFromAmount(0),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.b.Mul(tt.arg)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BPS.Mul() = %v, want %v", got, tt.want)
			}
			assertImmutableOperation(t, "BPS.Mul()", got, tt.b)
		})
	}
}

func TestBPS_Div(t *testing.T) {
	tests := []struct {
		name string
		b    *bps.BPS
		arg  int64
		want *bps.BPS
	}{
		{
			"100 ppms / 4 = 25 ppms",
			bps.NewFromPPM(big.NewInt(100)),
			4,
			bps.NewFromPPM(big.NewInt(25)),
		},
		{
			"100 ppms / 3 = 33 ppms, truncates after the decimal point",
			bps.NewFromPPM(big.NewInt(100)),
			3,
			bps.NewFromPPM(big.NewInt(33)),
		},
		{
			"100 ppms / -5 = -20 ppms",
			bps.NewFromPPM(big.NewInt(100)),
			-5,
			bps.NewFromPPM(big.NewInt(-20)),
		},
		{
			"nil / 5 = 0 ppms",
			&bps.BPS{},
			5,
			bps.NewFromPPM(big.NewInt(0)),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.b.Div(tt.arg)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BPS.Div() = %v, want %v", got, tt.want)
			}
			assertImmutableOperation(t, "BPS.Div()", got, tt.b)
		})
	}

	t.Run("If Div() by zero, it should occure division-by-zero run-time panic", func(t *testing.T) {
		b := bps.NewFromAmount(1)
		defer func() {
			err := recover()
			if err != "division by zero" {
				t.Errorf("got %s, want %s", err, "division by zero")
			}
		}()
		b.Div(0)
	})
}

func TestBPS_Compare(t *testing.T) {
	t.Run("zero and zero", func(t *testing.T) {
		t.Parallel()

		b := bps.NewFromAmount(0)
		b2 := bps.NewFromAmount(0)

		if got := b.Cmp(b2); got != 0 {
			t.Errorf("BPS.Cmp() = %d, want 0", got)
		}
		if !b.Equal(b2) {
			t.Error("BPS.Equal() = true")
		}
		if !b.IsZero() {
			t.Error("BPS.IsZero() = true")
		}
	})

	t.Run("About two differrent values", func(t *testing.T) {
		t.Parallel()

		b := bps.NewFromAmount(1)
		b2 := bps.NewFromAmount(2)

		if got := b.Cmp(b2); got != -1 {
			t.Errorf("BPS.Comp() = %d, want -1", got)
		}
		if b.Equal(b2) {
			t.Error("BPS.Equal() = false")
		}
		if b.IsZero() {
			t.Error("BPS.IsZero() = false")
		}
	})

	t.Run("nil and nil", func(t *testing.T) {
		t.Parallel()

		b := &bps.BPS{}
		b2 := &bps.BPS{}

		if got := b.Cmp(b2); got != 0 {
			t.Errorf("BPS.Cmp() = %d, want 0", got)
		}
		if !b.Equal(b2) {
			t.Error("BPS.Equal() = true")
		}
		if !b.IsZero() {
			t.Error("BPS.IsZero() = true")
		}
	})
}

func TestSum(t *testing.T) {
	tests := []struct {
		name  string
		first *bps.BPS
		rest  []*bps.BPS
		want  *bps.BPS
	}{
		{
			"1 amount + 1 percentage + 1 basis point + 1 deci basis point + 1 ppm + empty + nil = 1010,111 ppms",
			bps.NewFromAmount(1),
			[]*bps.BPS{
				bps.NewFromPercentage(1),
				bps.NewFromBasisPoint(1),
				bps.NewFromDeciBasisPoint(1),
				bps.NewFromPPM(big.NewInt(1)),
				{},
				nil,
			},
			bps.NewFromPPM(big.NewInt(1010111)),
		},
		{
			"1 amount + (-1) percentage + (-1) basis point + (-1) deci basis point + (-1) ppm = 989,889 ppms",
			bps.NewFromAmount(1),
			[]*bps.BPS{
				bps.NewFromPercentage(-1),
				bps.NewFromBasisPoint(-1),
				bps.NewFromDeciBasisPoint(-1),
				bps.NewFromPPM(big.NewInt(-1)),
			},
			bps.NewFromPPM(big.NewInt(989889)),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := bps.Sum(tt.first, tt.rest...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBPS_Abs(t *testing.T) {
	tests := []struct {
		name string
		b    *bps.BPS
		want *bps.BPS
	}{
		{
			"If plus value, it should return the same value",
			bps.NewFromAmount(1),
			bps.NewFromAmount(1),
		},
		{
			"If minus value, it should return the plus value",
			bps.NewFromAmount(-1),
			bps.NewFromAmount(1),
		},
		{
			"If nil, it should retrun zero",
			&bps.BPS{},
			bps.NewFromAmount(0),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.b.Abs()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BPS.Abs() = %v, want %v", got, tt.want)
			}
			assertImmutableOperation(t, "BPS.Abs()", got, tt.b)
		})
	}
}

func TestBPS_Neg(t *testing.T) {
	tests := []struct {
		name string
		b    *bps.BPS
		want *bps.BPS
	}{
		{
			"If plus value, it should return the minus value",
			bps.NewFromAmount(1),
			bps.NewFromAmount(-1),
		},
		{
			"If minus value, it should return the plus value",
			bps.NewFromAmount(-1),
			bps.NewFromAmount(1),
		},
		{
			"If nil, it should return zero",
			&bps.BPS{},
			bps.NewFromAmount(0),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.b.Neg()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BPS.Neg() = %v, want %v", got, tt.want)
			}
			assertImmutableOperation(t, "BPS.Neg()", got, tt.b)
		})
	}
}

func TestAvg(t *testing.T) {
	tests := []struct {
		name  string
		first *bps.BPS
		rest  []*bps.BPS
		want  *bps.BPS
	}{
		{
			"the average of 50 basis points, 125 basis points, and 345 basis points is 17,333 ppms rounded off",
			bps.NewFromBasisPoint(50),
			[]*bps.BPS{
				bps.NewFromBasisPoint(125),
				bps.NewFromBasisPoint(345),
			},
			bps.NewFromPPM(big.NewInt(17333)),
		},
		{
			"the average of 3 percentages, 2 amounts, and 50 basis points is 6783,333 ppms rounded off",
			bps.NewFromPercentage(3),
			[]*bps.BPS{
				bps.NewFromAmount(2),
				bps.NewFromBasisPoint(50),
			},
			bps.NewFromPPM(big.NewInt(678333)),
		},
		{
			"the average of 3 deci basis points, 15 percentages, and -360 basis points is 3801 deci basis points",
			bps.NewFromDeciBasisPoint(3),
			[]*bps.BPS{
				bps.NewFromPercentage(15),
				bps.NewFromBasisPoint(-360),
			},
			bps.NewFromDeciBasisPoint(3801),
		},
		{
			"the average of 50 basis points, 125 basis points, and nil is 5,833 ppms rounded off",
			bps.NewFromBasisPoint(50),
			[]*bps.BPS{
				bps.NewFromBasisPoint(125),
				{},
			},
			bps.NewFromPPM(big.NewInt(5833)),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := bps.Avg(tt.first, tt.rest...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Avg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name  string
		first *bps.BPS
		rest  []*bps.BPS
		want  *bps.BPS
	}{
		{
			"the maximum value of 10 basis points, 99 deci basis points, nil, and 1001 ppms is 1001 ppms",
			bps.NewFromBasisPoint(10),
			[]*bps.BPS{
				bps.NewFromDeciBasisPoint(99),
				{},
				bps.NewFromPPM(big.NewInt(1001)),
			},
			bps.NewFromPPM(big.NewInt(1001)),
		},
		{
			"the maximum value of -10 basis points, -99 deci basis points, and -1001 ppms is -99 deci basis points",
			bps.NewFromBasisPoint(-10),
			[]*bps.BPS{
				bps.NewFromDeciBasisPoint(-99),
				bps.NewFromPPM(big.NewInt(-1001)),
			},
			bps.NewFromDeciBasisPoint(-99),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := bps.Max(tt.first, tt.rest...); !got.Equal(tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		name  string
		fisrt *bps.BPS
		rest  []*bps.BPS
		want  *bps.BPS
	}{
		{
			"the minimum value of 10 basis points, 99 deci basis points, nil, and 1001 ppms is 0 ppms",
			bps.NewFromBasisPoint(10),
			[]*bps.BPS{
				bps.NewFromDeciBasisPoint(99),
				{},
				bps.NewFromPPM(big.NewInt(1001)),
			},
			bps.NewFromAmount(0),
		},
		{
			"the minimum value of -10 basis points, -99 deci basis points, and -1001 ppms is -1001 deci basis points",
			bps.NewFromBasisPoint(-10),
			[]*bps.BPS{
				bps.NewFromDeciBasisPoint(-99),
				bps.NewFromPPM(big.NewInt(-1001)),
			},
			bps.NewFromPPM(big.NewInt(-1001)),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := bps.Min(tt.fisrt, tt.rest...); !got.Equal(tt.want) {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBPS_FloatString(t *testing.T) {
	tests := []struct {
		name string
		b    *bps.BPS
		prec int
		want string
	}{
		{
			"1 ppm presents `0` as string",
			bps.NewFromPPM(big.NewInt(1)),
			0,
			"0",
		},
		{
			"1 ppm presents `0.000001` as string",
			bps.NewFromPPM(big.NewInt(1)),
			6,
			"0.000001",
		},
		{
			"1 deci basis point presents `0.00001` as string",
			bps.NewFromDeciBasisPoint(1),
			5,
			"0.00001",
		},
		{
			"1 deci basis point presents `0.000010` as string",
			bps.NewFromDeciBasisPoint(1),
			6,
			"0.000010",
		},
		{
			"1 basis point presents `0.0001` as string",
			bps.NewFromBasisPoint(1),
			4,
			"0.0001",
		},
		{
			"1 percentage presents `0.01` as string",
			bps.NewFromPercentage(1),
			2,
			"0.01",
		},
		{
			"5 percentage presents `0.1` as string, rounded to nearest",
			bps.NewFromPercentage(5),
			1,
			"0.1",
		},
		{
			"1 amount presents `1` as string",
			bps.NewFromAmount(1),
			0,
			"1",
		},
		{
			"1 amount presents `1.0` as string",
			bps.NewFromAmount(1),
			1,
			"1.0",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.b.FloatString(tt.prec); got != tt.want {
				t.Errorf("BPS.FloatString() = %v, want %v", got, tt.want)
			}
		})
	}
}
