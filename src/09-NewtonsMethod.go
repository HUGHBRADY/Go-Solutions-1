// Go Problem 9 - Newton's Method for Square Roots
// Hugh Brady - G00338260 - 05/10/17

// Code adapted from Ian McLoughlin's whiteboard

package main

import (
	"fmt"
	"math"
)

// Newton's equation for finding a square root
func z_next(z float64, x float64) float64{
	return z - (((z * z) - x) / (2 * z))
}

func main() {
	x := 20.0
	var z float64
	
	for z = 1.0; z != z_next(z, x); z = z_next(z, x) {
		fmt.Printf("Current guess: %.5f\n", z)
	}

	//Prints the square root based off the mathematical equation
	fmt.Printf("\nThe square root of %.2f is %.5f \n\n", x, z)
	
	//Prints the exact root using the math import
	fmt.Printf("math.Sqrt returns: %.5f", math.Sqrt(x));
}