package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Here we see always same combination, that is connected to seed
	// So we need to update seed
	rand.Seed(100)

	guess := 10

	for n := 0; n != guess; {
		n = rand.Intn(guess + 1)
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}
