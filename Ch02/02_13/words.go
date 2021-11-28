// Challenge : Maps
package main

import (
	"fmt"
	"strings"
)

func main () {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`

	// print out how many times each word appears in the text
	// use strings.Fields to split to words and
	// strings.ToLower to convert to lowercase

	// initialize the map containing the counter of reach word
	var wordCount = make(map[string]int, 20);
	//var wordCount = map[string]int{}
	// split words
	var words = strings.Fields(text);
	// for each word in the text
	for _,word := range words {
		// ignore case
		word = strings.ToLower(word);
		// check if have already counted this word
		if _, ok := wordCount[word]; ok {
			wordCount[word]++; // increment counter
		} else {
			wordCount[word] = 1; // initialize counter at 1
		}
	}

	fmt.Println(wordCount);

}
