package bps

import (
	"database/sql"
	"database/sql/driver"
	"encoding"
	"errors"
)

// make sure that the *BPS implements some interfaces.
var _ interface {
	sql.Scanner
	driver.Valuer
	encoding.TextUnmarshaler
} = (*BPS)(nil)

// Scan implements the sql.Scanner interface for database deserialization.
func (b *BPS) Scan(value interface{}) error {
	if b == nil {
		return errors.New("BPS.Scan: nil receiver")
	}

	switch v := value.(type) {
	case uint:
		s := NewFromBaseUnit(int64(v))
		b.value = s.value
		return nil
	case uint32:
		s := NewFromBaseUnit(int64(v))
		b.value = s.value
		return nil
	case uint64:
		s := NewFromBaseUnit(int64(v))
		b.value = s.value
		return nil
	case int:
		s := NewFromBaseUnit(int64(v))
		b.value = s.value
		return nil
	case int32:
		s := NewFromBaseUnit(int64(v))
		b.value = s.value
		return nil
	case int64:
		s := NewFromBaseUnit(v)
		b.value = s.value
		return nil
	case string:
		s, err := NewFromString(v)
		if err != nil {
			return err
		}
		b.value = s.value
		return nil
	}

	return errors.New("BPS.Scan: invalid type, supporting only integer or string")
}

// Value implements the driver.Valuer interface for database serialization.
func (b *BPS) Value() (driver.Value, error) {
	return b.String(), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (b *BPS) UnmarshalText(text []byte) error {
	v := string(text)
	if v == "" {
		return errors.New("BPS.UnmarshalText: no data")
	}

	n, err := NewFromString(v)
	if err != nil {
		return err
	}
	b.value = n.value

	return nil
}
