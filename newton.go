//Damian Gavin 25/09/17
//Write a function to calculate the square root of a number
//using Newtons method
//aided with golang playground https://tour.golang.org/flowcontrol/8

package main

//import math for comparissons
import (
	"fmt"
	"math"
)

const Delta = 0.0000001
const InitialZ = 600.0

//This is Newtons method, he approximates Sqrt by picking the starting
//point z and then repeating: z-(z*z - x)/(2*z)
func NewtonsSqrt(x float64) (z float64) {
	const IniitialZ = 100.0
	z = InitialZ

	step := func() float64 {
		return z - (z*z-x)/(2*z)
	}
	for zz := step(); math.Abs(zz-z) > Delta; {
		z = zz
		zz = step()
	}
	return
}

func main() {
	fmt.Println(NewtonsSqrt(500))
	fmt.Println(math.Sqrt(500)) //callled from math package
}
