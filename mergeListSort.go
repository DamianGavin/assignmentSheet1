//Damian Gavin 23/09/17
//Write a function that merges two sorted lists into a new sorted list.
// [1,4,6],[2,3,5] → [1,2,3,4,5,6].
//I have decided to use lists of Integers for this excercise

package main

import (
	"fmt"
	"sort"
)

func DoAppendSlices() {
	myIntSlice := []int{1, 4, 6, 9, 22, 55, 28, -10000000000}
	myIntSlice = append(myIntSlice, []int{2, 3, 5, 11234}...) //...is a holding
	//place for elements in the argument, in this case myIntSlice

	//Package sort provides primitives for sorting slices
	//and user-defined collections.It works ideal here and can handle
	//any ammount of entries into the arrays
	sort.Ints(myIntSlice)
	fmt.Println(myIntSlice)
}
func main() {
	DoAppendSlices()
}
