package statistics

import (
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func TestReadDataLines(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "data.txt")
	content := " 10\n\n20 \n  \nabc\n\t30\t\n"

	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("failed writing test file: %v", err)
	}

	got, err := ReadDataLines(path)
	if err != nil {
		t.Fatalf("ReadDataLines returned error: %v", err)
	}

	want := []string{"10", "20", "abc", "30"}
	if !slices.Equal(got, want) {
		t.Fatalf("ReadDataLines mismatch: got %v want %v", got, want)
	}
}

func TestReadDataLinesMissingFile(t *testing.T) {
	t.Parallel()

	_, err := ReadDataLines(filepath.Join(t.TempDir(), "missing.txt"))
	if err == nil {
		t.Fatal("expected error for missing file, got nil")
	}
}

func TestParseInputMixedValues(t *testing.T) {
	t.Parallel()

	lines := []string{"10", "abc", "20"}
	gotNums, gotWarnings, err := ParseInput(lines)
	if err != nil {
		t.Fatalf("ParseInput returned unexpected error: %v", err)
	}

	wantNums := []int{10, 20}
	if !slices.Equal(gotNums, wantNums) {
		t.Fatalf("nums mismatch: got %v want %v", gotNums, wantNums)
	}

	wantWarnings := []string{"Skipping invalid number: abc"}
	if !slices.Equal(gotWarnings, wantWarnings) {
		t.Fatalf("warnings mismatch: got %v want %v", gotWarnings, wantWarnings)
	}
}

func TestParseInputAllInvalid(t *testing.T) {
	t.Parallel()

	lines := []string{"bad", "nope"}
	gotNums, gotWarnings, err := ParseInput(lines)
	if err == nil {
		t.Fatal("expected error for all invalid input, got nil")
	}
	if gotNums != nil {
		t.Fatalf("expected nil nums, got %v", gotNums)
	}
	if len(gotWarnings) != 2 {
		t.Fatalf("expected 2 warnings, got %d", len(gotWarnings))
	}
}
