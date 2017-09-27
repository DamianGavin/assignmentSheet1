//Damian Gavin 14/09/2017
//This program will print the current date and time
//package required for all programmes
//Adapted from https://tour.golang.org/welcome/1
package main

import (
	"fmt"
	"time"
)

//I have formatted the time to look like the way
//we write it here!
//Just one function
func main() {
	t := time.Now()
	fmt.Println("Welcome to my 2nd go program")

	fmt.Printf("%-2d-%d-%4d Time:%02d:%02d\n",
		t.Day(), t.Month(), t.Year(),
		t.Hour(), t.Minute())
}
