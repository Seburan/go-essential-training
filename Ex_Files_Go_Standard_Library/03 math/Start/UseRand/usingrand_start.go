package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// TODO: Shuffle a sequence of values
	const numstring = "one two three four five six"
	words := strings.Fields(numstring);

	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i];
	})
	fmt.Println(words);


	// TODO: Generate a random string
	const upletters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const lowletters = "abcdefghijklmnopqrstuvwxyz"
	const digits = "0123456789"
	const specialchars = "~=+%^*()[]{}!@#$?|"
	const allchars = upletters + lowletters + digits + specialchars

	var sb strings.Builder;
	length := 10;

	for i:=0; i<length; i++ {
		sb.WriteRune(rune(allchars[rand.Intn(len(allchars))]));
	}

	fmt.Println("String Result :", sb.String());

	// TODO: Generate random string with rules
	stringBuf := make([]byte, length);

	stringBuf[0] = digits[rand.Intn(len(digits))];
	stringBuf[1] = specialchars[rand.Intn(len(specialchars))];
	stringBuf[2] = upletters[rand.Intn(len(upletters))];
	stringBuf[3] = lowletters[rand.Intn(len(lowletters))];

	for i:=4; i<length; i++ {
		stringBuf[i] = allchars[rand.Intn(len(allchars))];
	}

	rand.Shuffle(len(stringBuf), func(i, j int) {
		stringBuf[i], stringBuf[j] = stringBuf[j], stringBuf[i];
	})

	fmt.Println("Rules Result :", string(stringBuf));

}
