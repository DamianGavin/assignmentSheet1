//Damian Gavin 20/09/17
//program to return the largest and
//smallest element in a list.
//I have decided to use a list of integers for this problem.
//Adapted from https://gist.github.com/pyk/10519339

package main

import "fmt"

func main() {
	var n, smallest, biggest int
	x := []int{
		101, 296, -66, -429,
		45, 49, 63, 70,
		37, 509, 83, -27,
		91, 97, 99, 34,
		100, 1222223, 1, 3,
	}

	//range iterates over elements in a variety of data structures.
	for _, v := range x {
		if v > n {
			n = v
			biggest = n
		}

	}
	//can accomodate negative numbers also
	fmt.Printf("The biggest number is: %d\n", biggest)

	for _, v := range x {
		if v > n {
		} else {
			n = v
			smallest = n
		}
	}
	fmt.Printf("The smallest number is: %d ", smallest)
}
