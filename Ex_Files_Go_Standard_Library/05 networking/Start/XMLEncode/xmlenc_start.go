package main

import (
	"encoding/xml"
	"fmt"
)

type person struct {
	XMLName 	xml.Name	`xml:"persondata"`
	FirstName  string	`xml:"firstname"`
	LastName   string	`xml:"lastname"`
	Address    string	`xml:"address"`
	Age        int	`xml:"age,attr"`
	FaveColors []string	`xml:"favecolors"`
}

func main() {
	// TODO: Declare some sample data
	people := []person{
		{FirstName: "Jane", LastName: "Doe", Address: "123 Anywhere Street", Age: 35},
		{FirstName: "John", LastName: "Public", Address: "456 Everywhere Blvd", Age: 29,
			FaveColors: []string{"Purple", "Yellow", "Green"}},
	}

	//TODO: The MarhsalIndent function indents the XML text
	result, err := xml.MarshalIndent(people, "", "\t");
	if err != nil {
		fmt.Println("Error :", err);
	}
	fmt.Printf("%s%s", xml.Header, result);
}
