//Damian Gavin 22/06/17
//Program to find the factorial of 100
//and to sum the digits in the result
//adapted from https://play.golang.org/p/yW7sAyJpPe
package main

//math/big allows the use of large integers
import (
	"fmt"
	"math/big"
)

//main does the work, getting factorial for a range of
//numbers(100 in this case)

//calculates factorial of a number. Supports big numbers
func factorial(n int64) *big.Int {
	if n < 0 {
		return big.NewInt(1)
	}
	if n == 0 {
		return big.NewInt(1)
	}
	bigN := big.NewInt(n)
	return bigN.Mul(bigN, factorial(n-1))

}

//main calls factorial() and sums and prints the sum of digits
func main() {
	fmt.Println(factorial(100)) //100 can be changed to any number
	sum := big.NewInt(0)
	n := new(big.Int) //to store the factorial
	fact := (factorial(100))

	//This method will break up the answer into individual bits.
	for fact.BitLen() > 0 {
		_, n := fact.DivMod(fact, new(big.Int).SetUint64(uint64(10)), n)
		sum = sum.Add(sum, n)
	}
	//_, acts like a place holder. BitLen breaks the result into individual digits
	fmt.Println("\nThe sum of the digits in your answer is: ", (sum))
}
