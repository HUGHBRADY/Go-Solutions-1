// Go Problem 5 - Guessing Game
// Hugh Brady - G00338260 - 28/09/17

// randGen adapted from https://stackoverflow.com/questions/12321133/golang-random-number-generator-how-to-seed-properly

package main

import(
	"time"
	"fmt"
	"math/rand"
)

// function uses time as a seed to generate a random number
func randGen(min, max int) int{
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func main(){
	randNo:= randGen(1,30)
	var guess int
	var lastGuess int
	var counter int 
	
	fmt.Print("Enter a number between 1 and 30: ")

	// use counter as loop control
	for counter = 1; counter < 30; counter++{
		// prompt and read user input
		fmt.Scanf("%d", &guess)
		fmt.Scanf("%d")
		
		// if to check if guess is the same as the last
		if guess == lastGuess {
			fmt.Print("You just guessed that! Try again: ")
			counter--
		} else {
			// if the guess is a new one, then check the guess against the answer
			if guess == randNo {
				fmt.Print("Correctamundo! And it only took you ", counter, " guesses!\n")
				break
			} else if guess < randNo {
				fmt.Print("Wrong! Try a higher number: ")
			} else if guess > randNo {
				fmt.Print("Wrong! Try a lower number: ")
			}
		}// else if
		lastGuess = guess
	}// for
	
	if counter == 29 {
		// Let's you know you lose if you manage to enter everything but the answer
		fmt.Print("\nToo took to long! You lose!")
	}
	
}// main