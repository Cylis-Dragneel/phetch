package main

import (
	"bytes"
	"io"
	"os"
	"runtime"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	// Redirect stdout to capture output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call main function
	main()

	// Restore stdout
	w.Close()
	os.Stdout = old

	// Read captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Check if output contains expected information
	expectedStrings := []string{
		"System Information:",
		"OS: " + runtime.GOOS,
		"Architecture: " + runtime.GOARCH,
	}

	for _, str := range expectedStrings {
		if !strings.Contains(output, str) {
			t.Errorf("Expected output to contain '%s', but it didn't.\nActual output:\n%s", str, output)
		}
	}
}
