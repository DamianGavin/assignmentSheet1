//Damian Gavin 25/9/17
//This program will reverse a string in Go
//Date: 25.9.2017
//adapted from http://golangcookbook.com/chapters/strings/reverse/

package main

import "fmt"

/*Reversing a string by word is very similar to the process to reverse an array.
The difference is that strings are immutable in Go. Therefor, we must first
convert the string to a mutable array of runes ([]rune),
 perform the reverse operation on that, and then re-cast to a string.*/

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func main() {
	fmt.Printf("%v\n", reverse("damian gavin"))
}
