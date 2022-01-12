package main

import (
	"fmt"
	"math"
)

func main() {
	x := 10.0

	// TODO: Absolute value
	fmt.Println(math.Abs(x), math.Abs(-x));
	// TODO: Pow (x^y) and Exp (e^x)
	fmt.Println("10^2 =", math.Pow(10.0, 2));
	fmt.Println("e=", math.E, "e^10 =", math.Exp(10));

	// TODO: Trigonometry and other functions
	fmt.Println(math.Cos(math.Pi));
	fmt.Println(math.Sin(math.Pi/2));
	fmt.Println(math.Tan(math.Pi/2));
	fmt.Println(math.Log(2));


	// TODO: Square and cube roots
	fmt.Println("Sqrt :", math.Sqrt(25));
	fmt.Println("Sqrt :", math.Cbrt(125));

	// TODO: calculate the Hypotenuse of a right triangle
	fmt.Println("Hypothenus :", math.Hypot(30,40));

}
