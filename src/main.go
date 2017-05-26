package main

import (
	"fmt"
	"math/rand"
	"time"
)

func makeRand(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func main() {
	secretNumber := makeRand(1, 100)
	var theUsersGuess int
	chances := 7

	fmt.Println("GUESS THE NUMBER! (1-100)")

	for {
		if chances == 0 {
			fmt.Println("___ YOU LOOSE! ___")
			break
		}

		fmt.Println("Please input your guess | Chance(s): ", chances)
		if _, err := fmt.Scan(&theUsersGuess); err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		if theUsersGuess == secretNumber {
			fmt.Printf("___ YOU WIN, %d, Your score's: %d ___", secretNumber, chances*10)
			break
		} else if theUsersGuess < secretNumber {
			fmt.Println("Too small!")
		} else if theUsersGuess > secretNumber {
			fmt.Println("Too big!")
		}

		chances--
	}
}
