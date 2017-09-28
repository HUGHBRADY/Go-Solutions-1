// Go Problem 3 - Fizzbuzz
// Hugh Brady - G00338260 - 28/09/17

package main

import (
	"fmt"
)

func main() {
	var i int = 0
	
	// for loop running through numbers 1 - 100
	for i < 100{
		i++
		
		// check if i is divisible by 3 or 5
		if i%3 == 0 && i%5 == 0{
			fmt.Printf("FizzBuzz\n")
		} else {
			fmt.Printf("%d\n", i)
		}// end if else
			
		
		
	}
	
}