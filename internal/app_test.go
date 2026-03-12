package internal

import "testing"

func TestRun_H(t *testing.T) {
	code := Run([]string{"h", "9"})
	if code != 0 {
		t.Errorf("run h 9: exit code = %d, want 0", code)
	}
}

func TestRun_H_InvalidTz(t *testing.T) {
	code := Run([]string{"h", "-f", "Invalid/TZ", "9"})
	if code != 1 {
		t.Errorf("run h with invalid tz: exit code = %d, want 1", code)
	}
}

func TestRun_Dt(t *testing.T) {
	code := Run([]string{"dt", "2024-01-15T09:00:00"})
	if code != 0 {
		t.Errorf("run dt: exit code = %d, want 0", code)
	}
}

func TestRun_Unknown(t *testing.T) {
	code := Run([]string{"unknown"})
	if code != 1 {
		t.Errorf("run unknown: exit code = %d, want 1", code)
	}
}

func TestRun_NoArgs(t *testing.T) {
	code := Run([]string{})
	if code != 1 {
		t.Errorf("run no args: exit code = %d, want 1", code)
	}
}

func TestNewZoneFlags(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantFrom string
		wantTo   string
		wantArgs int
	}{
		{"defaults", []string{"9"}, "Asia/Tokyo", "UTC", 1},
		{"long flags", []string{"-from", "UTC", "-to", "Asia/Tokyo", "9"}, "UTC", "Asia/Tokyo", 1},
		{"short flags", []string{"-f", "UTC", "-t", "Asia/Tokyo", "9"}, "UTC", "Asia/Tokyo", 1},
		{"equals syntax", []string{"-from=UTC", "-to=Asia/Tokyo", "9"}, "UTC", "Asia/Tokyo", 1},
		{"no positional", []string{"-f", "UTC"}, "UTC", "UTC", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs, from, to := NewZoneFlags("test")
			if err := fs.Parse(tt.args); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if *from != tt.wantFrom {
				t.Errorf("from = %q, want %q", *from, tt.wantFrom)
			}
			if *to != tt.wantTo {
				t.Errorf("to = %q, want %q", *to, tt.wantTo)
			}
			if fs.NArg() != tt.wantArgs {
				t.Errorf("NArg() = %d, want %d", fs.NArg(), tt.wantArgs)
			}
		})
	}
}
