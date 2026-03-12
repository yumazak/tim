package internal

import (
	"testing"
	"time"
)

func TestConvertHour_JSTtoUTC(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	utc, _ := time.LoadLocation("UTC")

	tests := []struct {
		hour int
		want int
	}{
		{0, 15},
		{9, 0},
		{23, 14},
		{12, 3},
	}

	for _, tt := range tests {
		got := convertHour(tt.hour, jst, utc)
		if got != tt.want {
			t.Errorf("convertHour(%d, JST, UTC) = %d, want %d", tt.hour, got, tt.want)
		}
	}
}

func TestConvertHour_UTCtoJST(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	utc, _ := time.LoadLocation("UTC")

	tests := []struct {
		hour int
		want int
	}{
		{0, 9},
		{15, 0},
		{14, 23},
	}

	for _, tt := range tests {
		got := convertHour(tt.hour, utc, jst)
		if got != tt.want {
			t.Errorf("convertHour(%d, UTC, JST) = %d, want %d", tt.hour, got, tt.want)
		}
	}
}

func TestProcessHour_Valid(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	utc, _ := time.LoadLocation("UTC")

	result, err := processHour("9", jst, utc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "0" {
		t.Errorf("processHour(9) = %q, want %q", result, "0")
	}
}

func TestProcessHour_Invalid(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	utc, _ := time.LoadLocation("UTC")

	tests := []string{"24", "-1", "abc", ""}

	for _, input := range tests {
		_, err := processHour(input, jst, utc)
		if err == nil {
			t.Errorf("processHour(%q) expected error, got nil", input)
		}
	}
}

func TestFixedOffsetSeconds(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	off := fixedOffsetSeconds(jst)
	if off != 9*3600 {
		t.Errorf("fixedOffsetSeconds(JST) = %d, want %d", off, 9*3600)
	}

	utc, _ := time.LoadLocation("UTC")
	off = fixedOffsetSeconds(utc)
	if off != 0 {
		t.Errorf("fixedOffsetSeconds(UTC) = %d, want 0", off)
	}
}
