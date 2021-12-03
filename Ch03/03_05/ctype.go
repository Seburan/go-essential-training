// Challenge : Write a function

// Let's write a function that gets a URL and returns the value of
// Content-Type response HTTP header

// The function should return an error if it cannot perform a GET Request
package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println(`
	Let's write a function that gets a URL and returns the value of
	Content-Type response HTTP header
	`);

	var url = "https://golang.org";

	ctype, err := contentType(url);
	if err != nil {
		fmt.Printf("ERROR : %s\n", err)
	}	else {
		fmt.Printf("Content-Type of %s is %s\n", url, ctype );
	}

}

// Get content-type of fetched from the url passed as an argument
func contentType(url string) (string,error) {

	var ctype string;

	resp, err := http.Get(url);
	if err != nil {
		return "", err;
	}

	defer resp.Body.Close(); // Make sure we close the response body

	ctype = resp.Header.Get("Content-Type")
	if ctype == "" {
		return "", fmt.Errorf("cannot find content-type header from url");
	}

	return ctype, nil;
}
