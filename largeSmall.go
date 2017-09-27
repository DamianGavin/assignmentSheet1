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
		48, 96, 86, -68,
		57, 82, 63, 70,
		37, 34, 83, -27,
		19, 97, 9, 17,
		100, 1222223, 1, 3,
	}

	//range iterates over elements in a variety of data structures.
	for _, v := range x {
		if v > n {
			n = v
			biggest = n
		}

	}
	fmt.Println("The biggest number is ", biggest)

	for _, v := range x {
		if v > n {

			n = v
			smallest = n
		}
	}
	fmt.Println("The smallest number is ", smallest)
}
