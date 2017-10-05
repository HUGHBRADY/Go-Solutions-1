// Go Problem 5 - Largest and Smallest in a List
// Hugh Brady - G00338260 - 03/10/17

// code adapted from https://gist.github.com/pyk/10519339

package main

import(
	"fmt"
)

func main(){
	// variables and array
	var num, smallest, biggest int
	myList := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}
	
	// for loop to check biggest number
	for _,x:=range myList {
		if x>num {
			fmt.Println(x,">",num)
			
			num = x
			
			biggest = num
		} else {
			fmt.Println(x,"<",num)
		}
	}
		
	// for loop to check smallest number
	for _,x:=range myList {
		if x>num {
			fmt.Println(x,">",num)
		} else {
			fmt.Println(x,"<",num)
			
			num = x
			
			smallest = num
		}
	}
	
	// Print your findings
	fmt.Println("The biggest number is", biggest)
	fmt.Println("The smallest number is", smallest)
	
}// main