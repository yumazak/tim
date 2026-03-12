package internal

import (
	"testing"
	"time"
)

func TestProcessDatetime_Naive(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	utc, _ := time.LoadLocation("UTC")

	tests := []struct {
		input string
		want  string
	}{
		{"2024-01-15T09:00:00", "2024-01-15T00:00:00Z"},
		{"2024-01-15 09:00:00", "2024-01-15T00:00:00Z"},
		{"2024-06-15T12:30:00", "2024-06-15T03:30:00Z"},
		{"2024-01-15T09:00", "2024-01-15T00:00:00Z"},
		{"2024-01-15 09:00", "2024-01-15T00:00:00Z"},
	}

	for _, tt := range tests {
		got, err := processDatetime(tt.input, jst, utc)
		if err != nil {
			t.Fatalf("processDatetime(%q) unexpected error: %v", tt.input, err)
		}
		if got != tt.want {
			t.Errorf("processDatetime(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestProcessDatetime_WithOffset(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	utc, _ := time.LoadLocation("UTC")

	got, err := processDatetime("2024-01-15T09:00:00+09:00", jst, utc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "2024-01-15T00:00:00Z" {
		t.Errorf("got %q, want %q", got, "2024-01-15T00:00:00Z")
	}
}

func TestProcessDatetime_UTCtoJST(t *testing.T) {
	utc, _ := time.LoadLocation("UTC")
	jst, _ := time.LoadLocation("Asia/Tokyo")

	got, err := processDatetime("2024-01-15T00:00:00", utc, jst)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "2024-01-15T09:00:00+09:00" {
		t.Errorf("got %q, want %q", got, "2024-01-15T09:00:00+09:00")
	}
}

func TestProcessDatetime_Invalid(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	utc, _ := time.LoadLocation("UTC")

	tests := []string{"not-a-date", "2024-13-01T00:00:00", ""}

	for _, input := range tests {
		_, err := processDatetime(input, jst, utc)
		if err == nil {
			t.Errorf("processDatetime(%q) expected error, got nil", input)
		}
	}
}

func TestFormatDatetime_UTC(t *testing.T) {
	utc, _ := time.LoadLocation("UTC")
	dt := time.Date(2024, 1, 15, 0, 0, 0, 0, utc)
	got := formatDatetime(dt)
	if got != "2024-01-15T00:00:00Z" {
		t.Errorf("formatDatetime(UTC) = %q, want %q", got, "2024-01-15T00:00:00Z")
	}
}

func TestFormatDatetime_WithOffset(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	dt := time.Date(2024, 1, 15, 9, 0, 0, 0, jst)
	got := formatDatetime(dt)
	if got != "2024-01-15T09:00:00+09:00" {
		t.Errorf("formatDatetime(JST) = %q, want %q", got, "2024-01-15T09:00:00+09:00")
	}
}
