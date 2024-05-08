package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestGame(t *testing.T) {
	mockInput := "rock\npaper\nscissors\n0\n"
	reader := bufio.NewScanner(strings.NewReader(mockInput))

	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)

	Game(reader, writer)

	expectedOutput := "Computer's choice: rock \n" +
		"Computer's choice: paper \n" +
		"Computer's choice: scissors \n" +
		"Computer's choice: rock \n" +
		"Enter your choice (rock, paper, scissors, or 0 to exit): Exiting the game...\n"
	if buf.String() != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, buf.String())
	}
}

func Game(reader *bufio.Scanner, writer *bufio.Writer) {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	choices := [...]string{"rock", "paper", "scissors"}

	for {
		computerChoice := choices[rand.Intn(len(choices))]
		fmt.Fprintf(writer, "Computer's choice: %s \n", computerChoice)

		fmt.Fprint(writer, "Enter your choice (rock, paper, scissors, or 0 to exit): ")
		reader.Scan()
		userChoice := strings.ToLower(reader.Text())

		if userChoice == "0" {
			fmt.Fprintln(writer, "Exiting the game...")
			writer.Flush() // Flush the writer to ensure output is written
			break
		}

		if userChoice == computerChoice {
			fmt.Fprintln(writer, "It's a tie!")
		} else if (userChoice == "rock" && computerChoice == "scissors") ||
			(userChoice == "paper" && computerChoice == "rock") ||
			(userChoice == "scissors" && computerChoice == "paper") {
			fmt.Fprintln(writer, "You win!")
		} else {
			fmt.Fprintln(writer, "Computer wins!")
		}

		writer.Flush() // Flush the writer after each iteration
	}
}
