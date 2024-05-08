package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Define an array of choices
	choices := [3]string{"rock", "paper", "scissors"}

	// Create a new scanner for reading user input
	scanner := bufio.NewScanner(os.Stdin)

	// Game loop
	for {
		// Generate a random selection for computer
		randomIndex := rand.Intn(3)
		computerChoice := choices[randomIndex]
		fmt.Printf("Computer's choice: %s\n", computerChoice)

		// Prompt the user for their choice
		fmt.Print("Enter your choice (rock, paper, scissors, or 0 to exit): ")
		scanner.Scan()
		userChoice := strings.ToLower(scanner.Text())

		// Check if the user wants to exit
		if userChoice == "0" {
			fmt.Println("Exiting the game...")
			break
		}

		// Determine the winner
		if userChoice == computerChoice {
			fmt.Println("It's a tie!")
		} else if (userChoice == "rock" && computerChoice == "scissors") ||
			(userChoice == "paper" && computerChoice == "rock") ||
			(userChoice == "scissors" && computerChoice == "paper") {
			fmt.Println("You win!")
		} else if userChoice != "rock" && userChoice != "paper" && userChoice != "scissors" {
			fmt.Println("Invalid choice. Please try again.")
		} else {
			fmt.Println("Computer wins!")
		}
	}
}
