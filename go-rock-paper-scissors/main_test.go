// main_test.go
package main

import (
	"bufio"
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestGame(t *testing.T) {
	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)

	// Mock user input
	mockInput := "rock\npaper\nscissors\n0\n"
	reader := bufio.NewScanner(strings.NewReader(mockInput))

	// Redirect standard input and output
	oldReader := os.Stdin
	oldWriter := os.Stdout
	defer func() {
		os.Stdin = oldReader
		os.Stdout = oldWriter
	}()
	os.Stdin = os.NewFile(uintptr(0), "/dev/stdin")
	os.Stdout = os.NewFile(uintptr(1), "/dev/stdout")

	// Run the main function
	main()

	expectedOutput := "Computer's choice: rock\n" +
		"Computer's choice: paper\n" +
		"Computer's choice: scissors\n" +
		"Computer's choice: rock\n" +
		"Enter your choice (rock, paper, scissors, or 0 to exit): Exiting the game...\n"
	actualOutput := buf.String()

	if actualOutput != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, actualOutput)
	}
}
