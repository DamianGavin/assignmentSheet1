//Damian Gavin 20/09/17
//program to guess a secret number
//and count attempts
//REF:https://www.youtube.com/watch?v=HfJNvH9i9jI

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//this function will verify that an integer is entered
func askInt(message string) int {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("%s\n>>", message)
		scanner.Scan()
		s := scanner.Text()
		n, err := strconv.Atoi(s)

		if err == nil {
			return n
		}
		fmt.Printf("Error:'%s' is not a number\n", s)
	}
}

func guessIt() {
	//seed the number or it will always be the same
	rand.Seed(time.Now().UnixNano())
	var answer = rand.Intn(100)
	//counter for guesses, starts at 1 as they could be right 1st time!
	i := 1
	//prevG used to track if number was guessed before
	prevG := 0
	fmt.Println("I have a number")

	for {
		g := askInt("Take a guess")
		if g == prevG {
			fmt.Println("Sorry, you have already tried that number")
			i--
		}
		if g == answer {
			fmt.Println("You got it!")
			fmt.Printf("\nThe number of attempts was %d", i)
			return
		}
		if g < answer {
			fmt.Println("Nope: Higher")
		} else {
			fmt.Println("Nope: Lower")
		}
		i++
		prevG = g

	}

}

func main() {
	guessIt()
}
