package main

import (
	"fmt"
	"strings"
)

func main() {
	fname := "filename.txt"
	fname2 := "temp_picfile.jpeg"
	vowels := "aeiouAEIOU"
	s := "The quick brown fox jumps over the lazy dog"
	s2 := "aaaaaaaaaa"

	// Common string searching functions

	// Use Contains to see if a substring is in a string
	fmt.Println("Contains ? ", strings.Contains(s, "jumps"));

	// Use ContainsAny to see if any of the given chars are in the string
	fmt.Println("Contains any ? ", strings.ContainsAny(s, "abcdz"))

	// Use Index to find the offset of the first instance of a substring
	fmt.Println("Index of fox ?", strings.Index(s, "fox"));

	// Use IndexAny to find the first instance of any of the given chars
	fmt.Println("IndexAny of 'abcde' ?", strings.IndexAny(s, "abcde"));
	fmt.Println("IndexAny of vowels ?", strings.IndexAny("grzl", vowels));

	// HasPrefix and HasSuffix can be used to see if a string starts with
	// or ends with a specific substring
	fmt.Println("HasSuffix txt ?", strings.HasSuffix(fname, "txt"));
	fmt.Println("HasSuffix txt ?", strings.HasPrefix(fname2, "temp"));

	// Count returns the number of non-overlapping instances of a substring
	fmt.Println("Count 'the' occurrences : ", strings.Count(s, "the"));
	fmt.Println("Count 'he' occurrences : ", strings.Count(s, "he"));
	fmt.Println("Count 'aas' occurrences : ", strings.Count(s2, "aaa"));
}
