package main

	import "fmt"

	func main() {
		
	fmt.Println("starting fizzbuzz")
	c := make([]int, 100)
	for i := range c {
		d := i + 1
		threes := d%3 == 0
		fives := d%5 == 0
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