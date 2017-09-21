//Damian Gavin 20/09/17
//program to guess a secret number

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
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

func guessIt() {
	//seed the number or it will always be the same
	rand.Seed(time.Now().UnixNano())
	var answer = rand.Intn(100)
	//counter for guesses, starts at 1 as they could be right 1st time!
	i := 1

	fmt.Println("I'm thinking of a number")
	for {
		g := askInt("Take a guess")
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

	}

}

func main() {
	guessIt()
}
