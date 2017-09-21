//Damian Gavin 20/09/17
//program to guess a secret number

//Damian Gavin 14/09/17

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

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

func askYesHigherLower(message string) int {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("%s\n>>", message)
		scanner.Scan()
		s := scanner.Text()

		if s == "yes" {
			return 0
		}
		if s == "higher" {
			return 1
		}
		if s == "lower" {
			return -1
		}
		fmt.Printf("Please only enter, 'yes', 'higher', or 'lower'\n")
	}
}

func printToN(base int, top int) {
	for i := base; i <= top; i++ {
		fmt.Println(i)
	}
}

func keeper() {

	var answer int = rand.Intn(100)

	fmt.Println("I'm thinking of a number")
	for {
		g := askInt("Take a guess")
		if g == answer {
			fmt.Println("You got it!")
			return
		}
		if g < answer {
			fmt.Println("Nope: Higher")
		} else {
			fmt.Println("Nope: Lower")
		}
	}
}

//ask the user to think of a number
func seeker() {
	askYesHigherLower("Say 'yes' when you think of a number between 1 and 100")
	//loop to keep guessing until we find the answer

	highest := 100
	lowest := 1
	for guesses := 0; guesses < 10; guesses++ {

		mid := lowest
		if highest > lowest {
			mid = rand.Intn(highest-lowest) + lowest
		}
		res := askYesHigherLower("Is it " + strconv.Itoa(mid) +
			"? - yes, lower or higher?")
		if res == 0 {
			fmt.Println("Yay")
			return
		}
		if res > 0 {
			lowest = mid + 1
		}
		if res < 0 {
			highest = mid - 1
		}
		if lowest > highest {
			fmt.Printf("Don't be fucking cheating")
		}
	}
	fmt.Println("Too many guesses!")
}

func main() {
	//printToN(22,29)
	seeker()
	keeper()
}
