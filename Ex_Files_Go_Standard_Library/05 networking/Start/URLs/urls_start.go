package main

import (
	"fmt"
	"net/url"
)

func main() {
	// Define a URL
	s := "https://www.example.com:8000/user?username=joemarini"

	// TODO: Use the Parse function to parse the URL content
	result, err := url.Parse(s);
	if err != nil {
		panic(err);
	}
	fmt.Println(result.Scheme);
	fmt.Println(result.Host);
	fmt.Println(result.Port());
	fmt.Println(result.Path);
	fmt.Println(result.RawQuery);

	// TODO: Extract the query components into a Values struct
	values := result.Query();
	fmt.Println(values["username"]);

	// TODO: create a URL from components
	newUrl := &url.URL{
		Scheme: "https",
		Host: "www.google.com",
		Path: "/search",
		RawQuery: "q=golang",
	}

	s = newUrl.String();
	fmt.Println(s);

	// TODO: create a new Values struct and encode it as a query string
	newValues := url.Values{};
	newValues.Add("q", "python");
	newValues.Add("hl", "en");
	newUrl.RawQuery = newValues.Encode();

	s = newUrl.String();
	fmt.Println(s);
}
