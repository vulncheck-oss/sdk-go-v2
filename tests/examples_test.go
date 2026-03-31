package tests

import (
	"os/exec"
	"testing"
)

var examplePrograms = []string{
	"./examples/advisory.go",
	"./examples/backup.go",
	"./examples/backupv4.go",
	"./examples/connecting.go",
	"./examples/cpe.go",
	"./examples/index.go",
	"./examples/indices.go",
	"./examples/pagination.go",
	"./examples/purl.go",
}

func TestExamples(t *testing.T) {
	for _, program := range examplePrograms {
		program := program
		t.Run(program, func(t *testing.T) {
			cmd := exec.Command("go", "run", program)
			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("example failed with error: %v\n--- OUTPUT ---\n%s", err, out)
			}
		})
	}
}
