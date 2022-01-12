package main

import (
	"fmt"
	"strings"
)
func main() {
	// The map function returns a copy of a string with the characters modified
	// according to the mapping function
	shift := 4
	s := "The quick brown fox jumps over the lazy dog"

	// TODO: create the mapping function
	encoder := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			rPos := ((int(r) - int('A') + shift) % 26 ) + int('A')
			return rune(rPos);
		case r >= 'a' && r <= 'z':
			rPos := ((int(r) - int('a') + shift) % 26 ) + int('a')
			return rune(rPos);
		default:
			return r;
		}
	}
	decoder := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			rPos := ((int(r) - int('A') + 26 - shift) % 26 ) + int('A')
			return rune(rPos);
		case r >= 'a' && r <= 'z':
			rPos := ((int(r) - int('a') + 26 - shift) % 26 ) + int('a')
			return rune(rPos);
		default:
			return r;
		}
	}

	// TODO: Encode the message
	encodedString := strings.Map(encoder, s);
	fmt.Println(encodedString);

	// TODO: Decode the message
	decodedString := strings.Map(decoder, encodedString);
	fmt.Println(decodedString);

}
