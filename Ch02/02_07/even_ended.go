// Challenge : Even-ended numbers
package main

import (
	"fmt"
)

func main() {
	fmt.Println("How many even-ended numbers result from multiplying 2 four-digit numbers ?");

	var counter int;

	for i:=1000; i<10000; i++ {
		for j:=i; j<10000; j++ { // we do not want to count twice
			resultAsString := fmt.Sprintf("%d", i*j);
			// test if first digit = last digit
			if resultAsString[0] == resultAsString[len(resultAsString)-1] {
				// fmt.Printf("%v * %v = %v\n", i, j, resultAsString);
				counter++;
			}
		}
	}
	fmt.Printf("%v even-ended numbers result from multiplying 2 four-digit numbers\n", counter);
}
