package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

const maxTurn = 10

func main() {

	// To run the application, type : go run randomization.go <Some Number>

	arguments := os.Args

	// Here we check initial guess if provided
	if len(arguments) != 2 {
		fmt.Println("Initial guess value is not provided")
		return
	}

	// Here we check if initial guess is negative
	initialGuess, err := strconv.Atoi(arguments[1])

	if initialGuess < 0 {
		fmt.Println("Initial guess value should not be negative")
		return
	}

	// if there was some other error
	if err != nil {
		fmt.Printf("There was some error -> %v\n", err)
		return
	}

	for i := 0; i <= maxTurn; i++ {
		n := rand.Intn(initialGuess + 1)

		if n == initialGuess {
			fmt.Printf("You are Winner!!! Number of turns %d\n", i+1)
			return
		}
	}
	fmt.Println("You are Loser!!!!")

}
