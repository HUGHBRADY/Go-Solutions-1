// Go Problem 2 - Current Time
// Hugh Brady - G00338260 - 27/09/17

// Adapted from https://tour.golang.org/welcome/4

package main

import (
	"fmt"
	"time"		// "time" is used to measure and display time
)

// This function prints the current time
func main() {
	t := time.Now()
	fmt.Printf("The time is %d:%d", t.Hour(), t.Minute()) 
}
