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

func main() {
	fmt.Println(factorial(100))
	/* sum := 0
	for factorial.BitLen() > 0 {
		factorial.DivMod(factorial, new(big.Int).SetUint64(uint64(10)), n)
		sum += (n.Uint64())
	} */
}
