package bps

import "math/big"

// unit is the list of allowed values to set BaseUnit.
type unit int

// List of values that `unit` can take.
const (
	PPB unit = iota + 1
	PPM
	DeciBasisPoint
	HalfBasisPoint
	BasisPoint
	Percentage
)

// BaseUnit is unit to display *BPS as string via String method.
// Default is DeciBasisPoint unit, you can update this.
// But it should be used consistent value in your application.
var BaseUnit = DeciBasisPoint

type BPS struct {
	value *big.Int
}

// String returns the string representation of BaseUnit as generated by *big.Int.String().
// That means the effective digits is modifiable by BaseUnit.
func (b *BPS) String() string {
	return b.BaseUnitAmounts().String()
}
