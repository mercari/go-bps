package bps_test

import (
	"database/sql/driver"
	"reflect"
	"testing"

	"github.com/mercari/go-bps/bps"
)

func TestBPS_Value(t *testing.T) {
	tests := []struct {
		name string
		b    *bps.BPS
		want driver.Value
	}{
		{
			"zero",
			bps.NewFromAmount(0),
			"0",
		},
		{
			"1 amount",
			bps.NewFromAmount(1),
			"100000",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.b.Value()
			if err != nil {
				t.Errorf("BPS.Value() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BPS.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBPS_Scan(t *testing.T) {
	tests := []struct {
		name    string
		b       *bps.BPS
		value   interface{}
		want    *bps.BPS
		wantErr bool
	}{
		{
			"If b is nil, it should return an error",
			nil,
			"fake",
			nil,
			true,
		},
		{
			"If value is uint, it should set value as DeciBasisPoint",
			&bps.BPS{},
			uint(5),
			bps.NewFromDeciBasisPoint(5),
			false,
		},
		{
			"If value is uint32, it should set value as DeciBasisPoint",
			&bps.BPS{},
			uint32(6),
			bps.NewFromDeciBasisPoint(6),
			false,
		},
		{
			"If value is uint64, it should set value as DeciBasisPoint",
			&bps.BPS{},
			uint64(7),
			bps.NewFromDeciBasisPoint(7),
			false,
		},
		{
			"If value is int, it should set value as DeciBasisPoint",
			&bps.BPS{},
			int(6),
			bps.NewFromDeciBasisPoint(6),
			false,
		},
		{
			"If value is int32, it should set value as DeciBasisPoint",
			&bps.BPS{},
			int32(7),
			bps.NewFromDeciBasisPoint(7),
			false,
		},
		{
			"If value is int64, it should set value as DeciBasisPoint",
			&bps.BPS{},
			int64(8),
			bps.NewFromDeciBasisPoint(8),
			false,
		},
		{
			"If value is valid string, it should set value via NewFromString",
			&bps.BPS{},
			".15",
			bps.NewFromPercentage(15),
			false,
		},
		{
			"If value is invalid string, it should return an error",
			&bps.BPS{},
			"a15",
			&bps.BPS{},
			true,
		},
		{
			"If value is float, it should return an error",
			&bps.BPS{},
			.5,
			&bps.BPS{},
			true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := tt.b.Scan(tt.value); (err != nil) != tt.wantErr {
				t.Errorf("BPS.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.b, tt.want) {
				t.Errorf("BPS.Value() = %v, want %v", tt.b, tt.want)
			}
		})
	}
}
