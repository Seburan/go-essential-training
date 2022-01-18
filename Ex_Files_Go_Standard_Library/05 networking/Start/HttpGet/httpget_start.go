package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func getRequestTest() {
	const httpbin = "https://httpbin.org/get"

	// TODO: Perform a GET operation
	response, err := http.Get(httpbin);
	if err != nil {
		return;
	}

	// TODO: The caller is responsible for closing the response
	defer response.Body.Close();

	// TODO: We can access parts of the response to get information:
	fmt.Println("Status :", response.Status);
	fmt.Println("Status Code :", response.StatusCode);
	fmt.Println("Protocol :", response.Proto);
	fmt.Println("Content Length :", response.ContentLength);

	// TODO: Use a String Builder to build the content from bytes
	var sb strings.Builder;

	// TODO: Read the content and write it to the builder
	content, _ := ioutil.ReadAll(response.Body);
	byteCount, _ := sb.Write(content);

	// TODO: Format the output
	fmt.Println(byteCount, sb.String());

}

func main() {
	// Execute a GET request
	getRequestTest()
}
