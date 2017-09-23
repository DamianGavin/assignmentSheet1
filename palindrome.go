//Damian Gavin 23/09/17
//A function that tests whether a string is a palindrome.
//adapted from http://www.golangpro.com/2016/01/check-string-palindrome-go.html
package main

import (
	"fmt"
	"strings"
)

func main() {

	var ip string
	fmt.Println("Enter string:")
	fmt.Scanf("%s\n", &ip)
	ip = strings.ToLower(ip)
	fmt.Println(isP(ip))
}

//Function to test if the string entered is a Palindrome
//checks the outermost characters of the string and works back
//after finding the midpoint of string
//this method will accomodate numbers and spaces
func isP(s string) string {
	mid := len(s) / 2
	last := len(s) - 1
	for i := 0; i < mid; i++ {
		if s[i] != s[last-i] {
			return "No. That word or string is not a Palimdrome."
		}
	}
	return "Yes! You've entered a Palindrome"
}
