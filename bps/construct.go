package bps

import (
	"fmt"
	"math"
	"math/big"
	"strings"
)

// Denominators for each parts
const (
	DenomDeciBasisPoint int64 = 10
	DenomHalfBasisPoint       = DenomDeciBasisPoint * 5
	DenomBasisPoint           = DenomHalfBasisPoint * 2
	DenomPercentage           = DenomBasisPoint * 100
	DenomAmount               = DenomPercentage * 100
)

// NewFromString returns a new BPS from a string representation.
func NewFromString(value string) (*BPS, error) {
	var intString string
	var mul int64 = 1

	parts := strings.Split(value, ".")
	if len(parts) == 1 {
		// There is no decimal point, the original string can be just parsed with an int
		intString = value
	} else if len(parts) == 2 {
		// strip the insignificant digits for more accurate comparisons.
		decimalPart := strings.TrimRight(parts[1], "0")
		intString = parts[0] + decimalPart
		if intString == "" && parts[1] != "" {
			intString = "0"
		}
		expInt := len(decimalPart)
		mul = int64(math.Pow10(expInt))
	} else {
		return nil, fmt.Errorf("can't convert %s to BPS: too many .s", value)
	}

	parsed, ok := new(big.Int).SetString(intString, 10)
	if !ok {
		return nil, fmt.Errorf("can't convert %s to BPS", value)
	}

	return newBPS(parsed).Mul(DenomAmount).Div(mul), nil
}

// MustFromString returns a new BPS from a string representation or panics if NewFromString would have returned an error.
func MustFromString(value string) *BPS {
	b, err := NewFromString(value)
	if err != nil {
		panic(err)
	}
	return b
}

// NewFromPPM makes new BPS instance from part per million(ppm)
func NewFromPPM(ppm *big.Int) *BPS {
	return newBPS(ppm)
}

// NewFromDeciBasisPoint makes new BPS instance from deci basis point
func NewFromDeciBasisPoint(deci int64) *BPS {
	return newBPS(big.NewInt(deci)).Mul(DenomDeciBasisPoint)
}

// NewFromHalfBasisPoint makes new BPS instance from half basis point
func NewFromHalfBasisPoint(bp int64) *BPS {
	return newBPS(big.NewInt(bp)).Mul(DenomHalfBasisPoint)
}

// NewFromBasisPoint makes new BPS instance from basis point
func NewFromBasisPoint(bp int64) *BPS {
	return newBPS(big.NewInt(bp)).Mul(DenomBasisPoint)
}

// NewFromPercentage makes new BPS instance from percentage
func NewFromPercentage(per int64) *BPS {
	return newBPS(big.NewInt(per)).Mul(DenomPercentage)
}

// NewFromAmount makes new BPS instance from real amount
func NewFromAmount(amt int64) *BPS {
	return newBPS(big.NewInt(amt)).Mul(DenomAmount)
}

// NewFromBaseUnit makes new BPS instance from BaseUnit value.
// That means the effective digits is modifiable by BaseUnit.
func NewFromBaseUnit(v int64) *BPS {
	switch BaseUnit {
	case DeciBasisPoint:
		return NewFromDeciBasisPoint(v)
	case HalfBasisPoint:
		return NewFromHalfBasisPoint(v)
	case BasisPoint:
		return NewFromBasisPoint(v)
	case Percentage:
		return NewFromPercentage(v)
	}
	// The default unit is PPM
	return NewFromPPM(big.NewInt(v))
}

func newBPS(value *big.Int) *BPS {
	if value == nil {
		value = big.NewInt(0)
	}
	return &BPS{
		value: new(big.Int).Set(value),
	}
}
