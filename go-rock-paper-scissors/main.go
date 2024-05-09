package main

import (
	//	"bufio"
	"fmt"
	"math/rand"

	//	"os"
	"strings"
	"time"
)

func main() {
	//initialize random number generator
	theRandom := rand.New(rand.NewSource(time.Now().UnixNano()))

	//initialize valid choices
	choices := [3]string{"rock", "paper", "scissors"}

	var userWin, computerWin, tie int
	computerWin = 0
	userWin = 0
	tie = 0

	var computerChoice, userChoice string
	for {
		//computer choice
		computerChoice = choices[theRandom.Intn(3)]
		fmt.Printf("\n Computer's choice : %s\n", computerChoice)

		//player choice

		userChoice = getUserChoice()

		//result
		if userChoice == "0" {
			break
		} else if userChoice == computerChoice {
			fmt.Println("It's a tie!")
			tie++
		} else if (userChoice == "rock" && computerChoice == "scissors") ||
			(userChoice == "paper" && computerChoice == "rock") ||
			(userChoice == "scissors" && computerChoice == "paper") {
			fmt.Println("You win!")
			userWin++
		} else {
			fmt.Println("Computer wins!")
			computerWin++
		}

	}
	fmt.Println("Computer score = ", computerWin, "  User score = ", userWin, "  Tie = ", tie)
	if computerWin > userWin {
		fmt.Println("Computer wins!")
	} else if userWin > computerWin {
		fmt.Println("You win!")
	} else {
		fmt.Println("It's a tie!")
	}

}

func getUserChoice() string {
	var userMove string
	validChoices := map[string]bool{"rock": true, "paper": true, "scissors": true, "0": true}

	for {
		fmt.Print("Enter your choice (rock, paper, scissors, or 0 to exit): ")
		fmt.Scanln(&userMove)
		userMove = strings.ToLower(userMove)

		if _, ok := validChoices[userMove]; ok {
			break
		}

		fmt.Println("Invalid choice. Please try again.")
	}
	return userMove
}
