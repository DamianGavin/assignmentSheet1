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

	step := func() float64 { //makes step the actual mathamatical function
		return z - (z*z-x)/(2*z)
	}
	for zz := step(); math.Abs(zz-z) > Delta; { //abs is the absolute value
		z = zz
		zz = step()
	}
	return
}

func main() {
	fmt.Println("\nNewtons answer is:  ", (NewtonsSqrt(1500)))      //Newtons formula
	fmt.Println("\nThe Math package answer is:  ", math.Sqrt(1500)) //callled from math package
}
