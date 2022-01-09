package main

import "fmt"

type circle struct {
	radius int
	border int
}

func main() {
	x := 20
	f := 123.45

	// TODO: Basic formatting
	fmt.Printf("%d", x);
	fmt.Printf("%x\n", x);

	// TODO: Booleans can be printed as "true" or "false"
	fmt.Printf("%t\n", x > 10);

	// TODO: floating point numbers
	fmt.Printf("%f\n", f);
	fmt.Printf("%e\n", f);

	// TODO: Using explicit argument indexes
	fmt.Printf("%[2]d %[1]d\n", 45, 58)

	// TODO: Argument indexes can be used to print values repeatedly
	fmt.Printf("%[1]d %[1]b %[1]x\n", 32, 58)
	fmt.Printf("%#[1]d %#[1]b %#[1]x\n", 32, 58)

	// TODO: Print a value in default format
	c := circle{
		radius: 20,
		border: 5,
	}

	fmt.Printf("%v\n", c);
	fmt.Printf("%+v\n", c);
	fmt.Printf("%T\n", c);

	// TODO: Sprintf is the same as Printf, but returns a string
	s := fmt.Sprintf("%[2]d %[1]d\n", 45, 58)
	fmt.Println(s);
}
