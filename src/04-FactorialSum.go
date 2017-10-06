// Go Problem 4 - Factorial Digit Sum
// Hugh Brady - G00338260 - 28/09/17

// Code adapted from  https://play.golang.org/p/5lZMq3bd8a

package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Print("The Factorial of 100 is: ")
	fmt.Print(factorial(100))

	fmt.Print("\n\nThe sum of the Factorial of 100 is: ")
	fmt.Print(sumDigits(factorial(100)))
}

// Calculates factorial
func factorial(n int64) *big.Int {
   if n < 0 {
      return big.NewInt(1)
   }
   if (n==0) {
      return big.NewInt(1)
   }
   bigN := big.NewInt(n)
   return bigN.Mul(bigN, factorial(n-1))
}

// Calculates the sum of the digits in the factorial 
func sumDigits(number *big.Int) *big.Int {
    ten := big.NewInt(10)
    sum := big.NewInt(0)
    mod := big.NewInt(0)
    for ten.Cmp(number)<0 {
      sum.Add(sum, mod.Mod(number,ten))
      number.Div(number,ten)
    }	
	sum.Add(sum,number)
	return sum
}