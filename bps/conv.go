package bps

import "math/big"

// rawValue returns the row value as new big.Int instance
func (b *BPS) rawValue() *big.Int {
	s := nilSafe(b)
	return new(big.Int).Set(s.value)
}

// PPMs returns the row value that means PPM.
func (b *BPS) PPMs() *big.Int {
	return b.rawValue()
}

// Amounts returns the basis point as an integer amount.
func (b *BPS) Amounts() int64 {
	return b.Div(DenomAmount).rawValue().Int64()
}

// Percentages returns the basis point as an integer percentage count.
func (b *BPS) Percentages() *big.Int {
	return b.Div(DenomPercentage).rawValue()
}

// BasisPoints returns the basis point as an integer basis point count.
func (b *BPS) BasisPoints() *big.Int {
	return b.Div(DenomBasisPoint).rawValue()
}

// HalfBasisPoints returns the basis point as an integer half basis point count.
func (b *BPS) HalfBasisPoints() *big.Int {
	return b.Div(DenomHalfBasisPoint).rawValue()
}

// DeciBasisPoints returns the basis point as an integer half basis point count.
func (b *BPS) DeciBasisPoints() *big.Int {
	return b.Div(DenomDeciBasisPoint).rawValue()
}

// Rat returns a rational number representation of `b`.
func (b *BPS) Rat() *big.Rat {
	mul := big.NewInt(DenomAmount)
	num := nilSafe(b).value
	return new(big.Rat).SetFrac(num, mul)
}

// Float64 returns the nearest float64 value for `b` and a bool indicating whether f represents `b` exactly.
// If the magnitude of `b` is too large to be represented by a float64, f is an infinity and exact is false.
// The sign of f always matches the sign of `b`, even if f == 0.
func (b *BPS) Float64() (f float64, exact bool) {
	return b.Rat().Float64()
}

// BaseUnitAmounts returns amount representation of BaseUnit as generated.
// That means the effective digits is modifiable by BaseUnit.
func (b *BPS) BaseUnitAmounts() *big.Int {
	switch BaseUnit {
	case DeciBasisPoint:
		return b.DeciBasisPoints()
	case HalfBasisPoint:
		return b.HalfBasisPoints()
	case BasisPoint:
		return b.BasisPoints()
	case Percentage:
		return b.Percentages()
	}
	// default is PPM
	return b.rawValue()
}

// nilSafe returns zero value when b is nil or b.value is nil to avoid nil error.
func nilSafe(b *BPS) *BPS {
	if b == nil {
		return zero
	}
	if b.value == nil {
		return zero
	}
	return b
}
