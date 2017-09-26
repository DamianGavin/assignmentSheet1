//Damian Gavin 23/09/17
//Write a function that merges two sorted lists into a new sorted list.
// [1,4,6],[2,3,5] → [1,2,3,4,5,6].

package main

import "fmt"

func DoAppendSlices() {
	myIntSlice := []int{1, 4, 6}
	myIntSlice = append(myIntSlice, []int{2, 3, 5}...)
	fmt.Println(myIntSlice)
}
func main() {
	DoAppendSlices()
}
