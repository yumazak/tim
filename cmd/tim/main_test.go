package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestIntegration_H(t *testing.T) {
	binary := buildBinary(t)

	cmd := exec.Command(binary, "h", "9")
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("command failed: %v", err)
	}
	if strings.TrimSpace(string(out)) != "0" {
		t.Errorf("got %q, want %q", strings.TrimSpace(string(out)), "0")
	}
}

func TestIntegration_H_Stdin(t *testing.T) {
	binary := buildBinary(t)

	cmd := exec.Command(binary, "h")
	cmd.Stdin = strings.NewReader("0\n9\n23\n")
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("command failed: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	expected := []string{"15", "0", "14"}
	if len(lines) != len(expected) {
		t.Fatalf("got %d lines, want %d", len(lines), len(expected))
	}
	for i, line := range lines {
		if line != expected[i] {
			t.Errorf("line %d: got %q, want %q", i, line, expected[i])
		}
	}
}

func TestIntegration_H_Stdin_PartialError(t *testing.T) {
	binary := buildBinary(t)

	cmd := exec.Command(binary, "h")
	cmd.Stdin = strings.NewReader("9\nabc\n12\n")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err == nil {
		t.Fatal("expected non-zero exit code")
	}

	lines := strings.Split(strings.TrimSpace(stdout.String()), "\n")
	if len(lines) != 2 {
		t.Errorf("expected 2 stdout lines, got %d: %v", len(lines), lines)
	}
	if !strings.Contains(stderr.String(), "invalid hour") {
		t.Errorf("expected stderr to contain 'invalid hour', got %q", stderr.String())
	}
}

func TestIntegration_Dt(t *testing.T) {
	binary := buildBinary(t)

	cmd := exec.Command(binary, "dt", "2024-01-15T09:00:00")
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("command failed: %v", err)
	}
	if strings.TrimSpace(string(out)) != "2024-01-15T00:00:00Z" {
		t.Errorf("got %q, want %q", strings.TrimSpace(string(out)), "2024-01-15T00:00:00Z")
	}
}

func TestIntegration_Dt_WithFlags(t *testing.T) {
	binary := buildBinary(t)

	cmd := exec.Command(binary, "dt", "-f", "UTC", "-t", "Asia/Tokyo", "2024-01-15T00:00:00")
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("command failed: %v", err)
	}
	if strings.TrimSpace(string(out)) != "2024-01-15T09:00:00+09:00" {
		t.Errorf("got %q, want %q", strings.TrimSpace(string(out)), "2024-01-15T09:00:00+09:00")
	}
}

func buildBinary(t *testing.T) string {
	t.Helper()
	binary := t.TempDir() + "/tim"
	cmd := exec.Command("go", "build", "-o", binary, "./cmd/tim")
	cmd.Dir = "/Users/yumazak/dev/personal/tim"
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("build failed: %v\n%s", err, out)
	}
	return binary
}
