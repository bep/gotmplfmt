package tmplfmt

import (
	"flag"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// To update the golden files, set writeOutput to true and run `go test -update`.
var (
	update = flag.Bool("update", false, "update the golden files")
	copy   = flag.Bool("copy", false, "copy golden files from in to out")
)

func TestGolden(t *testing.T) {
	if *update {
		t.Log("Updating golden files...")
	}

	goldenDir := "golden"
	goldenDirIn := filepath.Join(goldenDir, "in")
	goldenDirOut := filepath.Join(goldenDir, "out")

	// Read golden/in.
	if err := filepath.Walk(goldenDirIn, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		baseName := strings.TrimPrefix(path, goldenDirIn+string(os.PathSeparator))
		t.Run(path, func(t *testing.T) {
			input, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("failed to read input file: %v", err)
			}
			goldenPath := filepath.Join(goldenDirOut, baseName)
			if *copy {
				if err := os.WriteFile(goldenPath, input, 0o644); err != nil {
					t.Fatalf("failed to copy golden file: %v", err)
				}
				t.Logf("Copied %s to %s", path, goldenPath)
				// Done
				return
			}
			output, err := Format(string(input))
			if err != nil {
				t.Fatalf("failed to format template: %v", err)
			}

			if *update {
				if err := os.WriteFile(goldenPath, []byte(output), 0o644); err != nil {
					t.Fatalf("failed to write golden file: %v", err)
				}
			} else {
				expected, err := os.ReadFile(goldenPath)
				if err != nil {
					t.Fatalf("failed to read golden file: %v", err)
				}
				if output != string(expected) {
					t.Errorf("output does not match golden file.\nGot:\n%s\nExpected:\n%s", output, expected)
				}
			}
		})
		return nil
	}); err != nil {
		t.Fatalf("failed to walk golden/in: %v", err)
	}
}
