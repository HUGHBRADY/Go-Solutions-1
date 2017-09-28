// Go Problem 5 - Guessing Game
// Hugh Brady - G00338260 - 28/09/17

package main

import (
	"time"
	"fmt"
	"math/rand"
)

func main() {
	rand.seed
	fmt.Print(rand.Intn(100))
}