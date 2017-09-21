//Damian Gavin FizzBuzz 19/09/17
//This program will print all integers from 1-100
//and will substitute any factor of 3 with "fizz" and
//any factor of 5 with "Buzz"

//Adapted from http://wiki.c2.com/?FizzBuzzTest

package main

//import format
	import "fmt"

	func main() {
		
	fmt.Println("Starting fizzbuzz in Go")
	c := make([]int, 100)
	for i := range c {
		d := i + 1
		threes := d%3 == 0//if divisible
		fives := d%5 == 0 //if divisible
		if threes && fives {
			fmt.Println("FizzBuzz")
		} else if threes {
			fmt.Println("Fizz")
		} else if fives {
			fmt.Println("Buzz")
		} else {
			fmt.Println(d)
		}
	}
	}