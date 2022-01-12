package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// declare sample strings
	s := "the quick brown fox jumped over the lazy dog"
	s2 := []string{"one", "two", "three", "four"}
	s3 := "This is a string. With some punctionation, for a demo! Yep."

	// The Split function splits a string into substrings
	sub1 := strings.Split(s, " ");
	fmt.Printf("%q\n", sub1);
	sub2 := strings.Split(s, "the");
	fmt.Printf("%q\n", sub2);

	// Join concatenates substrings, with the separator between each
	result := strings.Join(s2, " - ");
	fmt.Println(result);

	// The Fields function splits a string around white space
	sub3 := strings.Fields(s);
	fmt.Printf("%q\n", sub3);

	// FieldsFunc is a customizable version of fields that uses a callback
	f := func(r rune) bool {
		return unicode.IsPunct(r);
	}

	sub4 := strings.FieldsFunc(s3, f);
	fmt.Printf("%q\n", sub4);

	// Replacer can be used for multiple replacement operations
	rep := strings.NewReplacer(".","|",",","|","!","|");
	result2 := rep.Replace(s3);
	fmt.Println(result2);

}
