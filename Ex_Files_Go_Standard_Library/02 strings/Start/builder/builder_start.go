package main

import (
	"fmt"
	"strings"
)

func main() {
	// TODO: Declare an empty strings.Builder
	var	sb strings.Builder;

	// TODO: Write some content
	sb.WriteString("This is a string1\n");
	sb.WriteString("This is a string2\n");
	sb.WriteString("This is a string3\n");


	// TODO: Output the concatenated string
	fmt.Println(sb.String());

	// TODO: Examine the builder's capacity
	fmt.Println(sb.Cap());

	// TODO:
	// Grow the capacity - use this if you know in advance how much data
	sb.Grow(1024);
	fmt.Println("Capacity:", sb.Cap());

	// you're going to be writing into the buffer to minimize copies
	for i := 0; i < 10; i++ {
		fmt.Fprintf(&sb, "String %d --\n", i);
	}
	fmt.Println(sb.String());

	// TODO: we can get the length of what the final string will be
	fmt.Println("Length:", sb.Len());

	// TODO: The Reset function will reset the builder to original state
	sb.Reset();
	fmt.Println("Capacity:", sb.Cap());
	fmt.Println("Length:", sb.Len());
	fmt.Println(sb.String());

}
