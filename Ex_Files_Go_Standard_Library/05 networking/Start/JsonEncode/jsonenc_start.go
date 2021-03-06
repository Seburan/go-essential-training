package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name       string	`json:"fullname"`
	Address    string	`json:"addr"`
	Age        int	`json:"age"`
	FaveColors []string	`json:"favecolors,omitempty"`
}

func encodeExample() {
	// TODO: create some people data
	people := []person{
		{"Jane Doe", "123 Anywhere Street", 35, nil},
		{"John Public", "456 Everywhere Blvd", 31, []string{"Purple", "Yellow", "Green"}},
	}

	// TODO: Marshal is used to convert a data structure to JSON format
	// MarshalIndent is used to format the JSON string with indentation
	result, err := json.MarshalIndent(people, "", "\t");
	if err != nil {
		fmt.Println("Error: ", err);
	}
	fmt.Printf("%s\n", result);

}

func main() {
	// Encode Go data as JSON
	encodeExample()
}
