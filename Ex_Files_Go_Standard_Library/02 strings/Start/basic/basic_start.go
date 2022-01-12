package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "The quick brown fox jumps over the lazy dog"

	// Basic string operations
	fmt.Println(s);

	// TODO: Length of string
	fmt.Println("length : ", len(s));

	// TODO: iterate over each character
	for _, character := range(s) {
		fmt.Print(string(character), ",");
	}
	fmt.Println();

	// TODO: Using operators < > == !=
	fmt.Println("Dog" < "Cat");
	fmt.Println("Dog" < "Horse");
	fmt.Println("Dog" == "dog");

	// TODO: Comparing two strings
	result := strings.Compare("dog", "cat");
	fmt.Println(result);
	result = strings.Compare("dog", "dog");
	fmt.Println(result);

	// TODO: EqualFold tests using Unicode case-folding
	resultBool := strings.EqualFold("dog", "Dog");
	fmt.Println(resultBool);

	// TODO: ToUpper, ToLower, Title
	s1 := strings.ToUpper(s);
	s2 := strings.ToLower(s);
	s3 := strings.Title(s);
	fmt.Println(s1);
	fmt.Println(s2);
	fmt.Println(s3);
}
