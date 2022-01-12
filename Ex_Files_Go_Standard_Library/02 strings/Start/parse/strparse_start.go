package main

import (
	"fmt"
	"strconv"
)

func main() {
	sampleint := 100
	samplestr := "250"

	// This does a character conversion, not a numerical one
	// newstr := string(sampleint)
	// fmt.Println("Result of using string():", newstr)

	// The strconv package contains a variety of functions for parsing and formatting
	// numbers, values, and strings
	s := strconv.Itoa(sampleint);
	fmt.Printf("Itoa Type : %T - Valeue :%v\n", s, s);

	i, err := strconv.Atoi(samplestr);
	if err != nil {
		fmt.Println("Atoi error :", err.Error());
	}
	fmt.Printf("Atoi Type : %T - Valeue :%v\n", i, i);

	// Convert an integer to string

	// Convert a string to integer

	// Other parse functions
	b, _ := strconv.ParseBool("true");
	fmt.Printf("ParseBool Type : %T - Value :%v\n", b, b);
	f, _ := strconv.ParseFloat("3.14159", 64);
	fmt.Printf("ParrseFloat Type : %T - Value :%v\n", f, f);
	u, _ := strconv.ParseUint("12345", 10, 64);
	fmt.Printf("ParseUint Type : %T - Value :%v\n", u, u);

	// Other format functions
	s = strconv.FormatBool(b);
	fmt.Println("FormatBool:", s)
	s = strconv.FormatFloat(f, 'E', -1, 64 );
	fmt.Println("FormatFloat:", s)
	s = strconv.FormatUint(u, 10 );
	fmt.Println("FormatUin:", s)


}
