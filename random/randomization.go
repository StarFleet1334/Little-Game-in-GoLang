package main

import (
	"fmt"
	"math/rand"
	time2 "time"
)

func main() {
	// This way seed gets updated each time we call main
	// So we get different result each time
	time := time2.Now()
	rand.Seed(time.UnixNano())

	guess := 10

	for n := 0; n != guess; {
		n = rand.Intn(guess + 1)
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}
