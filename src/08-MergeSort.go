// Go Problem 8 - Merge List and Sort
// Hugh Brady - G00338260 - 05/10/17

// Code adapted from https://play.golang.org/p/g3QTWcy9D-

package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []int{
		99,15,48,11,45,16,
	}
	y := []int{
		65,10,71,24,53,5,
	}
	
	fmt.Printf("List X is: %v\n", x)
	fmt.Printf("List Y is: %v\n", y)

	mergedList:= merge(x, y)

	fmt.Printf("Merged list is: %v\n", mergedList)
}

func merge(s1, s2 []int) []int {
	// Create a map that holds the values from each slice
	unique := make(map[int]struct{}) // zero byte payload

	for _, v := range s1 {
		unique[v] = struct{}{} // zero byte payload
	}

	for _, v := range s2 {
		unique[v] = struct{}{} // zero byte payload
	}

	final := make([]int, len(unique)) // allocate to right size
	i := 0
	for k := range unique {
		final[i] = k
		i++ // index not a feature of map ranges
	}
	sort.Ints(final) // consistent order
	return final
}
