package bps

import "math/big"

// Abs returns the absolute value of the decimal.
func (b *BPS) Abs() *BPS {
	s := nilSafe(b)
	abs := new(big.Int).Abs(s.value)
	return newBPS(abs)
}

// Neg returns -b.
func (b *BPS) Neg() *BPS {
	s := nilSafe(b)
	neg := new(big.Int).Neg(s.value)
	return newBPS(neg)
}

// Add returns b + b2.
func (b *BPS) Add(b2 *BPS) *BPS {
	added := new(big.Int).Add(b.rawValue(), b2.rawValue())
	return newBPS(added)
}

// Sub returns b - b2.
func (b *BPS) Sub(b2 *BPS) *BPS {
	subbed := new(big.Int).Sub(b.rawValue(), b2.rawValue())
	return newBPS(subbed)
}

// Mul returns b * i.
func (b *BPS) Mul(i int64) *BPS {
	muled := new(big.Int).Mul(b.rawValue(), big.NewInt(i))
	return newBPS(muled)
}

// Div returns b / i, rounded down to ppm.
func (b *BPS) Div(i int64) *BPS {
	dived := new(big.Int).Div(b.rawValue(), big.NewInt(i))
	return newBPS(dived)
}

func (b *BPS) Cmp(b2 *BPS) int {
	return b.rawValue().Cmp(b2.rawValue())
}

func (b *BPS) Equal(b2 *BPS) bool {
	return b.Cmp(b2) == 0
}

var zero = &BPS{
	value: big.NewInt(0),
}

func (b *BPS) IsZero() bool {
	return b.Equal(zero)
}

// FloatString returns the string representation of the amount as generated by *big.Rat.FloatString(prec)
func (b *BPS) FloatString(prec int) string {
	return b.Rat().FloatString(prec)
}

// Avg returns the average value of the provided first and rest BPS
func Avg(first *BPS, rest ...*BPS) *BPS {
	count := int64(len(rest) + 1)
	sum := Sum(first, rest...)
	return sum.Div(count)
}

// Sum returns the combined total of the provided first and rest BPS
func Sum(first *BPS, rest ...*BPS) *BPS {
	total := first
	for _, b := range rest {
		total = total.Add(b)
	}
	return total
}

// Max returns the largest BPS that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//     Max(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call Max with 0 arguments.
func Max(first *BPS, rest ...*BPS) *BPS {
	max := first
	for _, b := range rest {
		if b.Cmp(max) > 0 {
			max = b
		}
	}
	return max
}

// Min returns the smallest BPS that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//     Min(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call Min with 0 arguments.
func Min(fisrt *BPS, rest ...*BPS) *BPS {
	min := fisrt
	for _, b := range rest {
		if b.Cmp(min) < 0 {
			min = b
		}
	}
	return min
}
