package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	choices := [...]string{"rock", "paper", "scissors"}

	for {
		computerChoice := choices[rand.Intn(len(choices))]
		fmt.Printf("Computer's choice: %s\n", computerChoice)

		userChoice := getUserChoice()

		if userChoice == "0" {
			fmt.Println("Exiting the game...")
			break
		}

		if userChoice == computerChoice {
			fmt.Println("It's a tie!")
		} else if (userChoice == "rock" && computerChoice == "scissors") ||
			(userChoice == "paper" && computerChoice == "rock") ||
			(userChoice == "scissors" && computerChoice == "paper") {
			fmt.Println("You win!")
		} else {
			fmt.Println("Computer wins!")
		}
	}
}

func getUserChoice() string {
	var userChoice string
	for {
		fmt.Print("Enter your choice (rock, paper, scissors, or 0 to exit): ")
		fmt.Scanln(&userChoice)
		userChoice = strings.ToLower(userChoice)

		if userChoice == "0" || userChoice == "rock" || userChoice == "paper" || userChoice == "scissors" {
			break
		}

		fmt.Println("Invalid choice. Please try again.")
	}
	return userChoice
}
