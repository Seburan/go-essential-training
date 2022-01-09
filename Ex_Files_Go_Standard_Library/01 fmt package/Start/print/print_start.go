package main

import "fmt"

func main() {
	// TODO: Print a simple string, no newline
  fmt.Print("Welcome to Go !")

	// TODO: Print with a newline
	fmt.Println("This ends with a new line")

	// TODO: Print a string with values
	const answer = 42
	length, err := fmt.Println("The answer to everything is", answer);
	fmt.Println(length, err);

	// TODO: Print a string with multiple interspersed values
	const a, b, c = 5, 5, 10
	fmt.Println("Add", a, "and", b, "you get", c);

	// TODO: print a slice of data
	items := []int{10, 20, 40, 80}
	length, err = fmt.Println(items);
	fmt.Println(length, err);

}
