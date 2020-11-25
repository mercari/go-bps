package bps

import (
	"database/sql/driver"
	"errors"
)

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
