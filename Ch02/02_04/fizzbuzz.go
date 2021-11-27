// Challenge : FizzBuzz
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Challenge : Fizzbuzz");

	// for every number from 1 to 99
	for i:=0; i < 100; i++ {
		switch {
		// if the number is divisible by 3 and by 5 print fizz buzz
		case i%3 == 0 && i%5 ==0:
			fmt.Println("fizz buzz");
		// if the number is divisible by 3 then print fizz buzz
		case i%3 == 0:
			fmt.Println("fizz");
		// if the number is divisible by 5 then print fizz buzz
		case i%5 == 0:
			fmt.Println("buzz");
		default:
			fmt.Println(i);
		}
	}
}
