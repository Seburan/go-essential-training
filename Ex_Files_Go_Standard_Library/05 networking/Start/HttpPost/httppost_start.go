package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func postRequestTest() {
	const httpbin = "https://httpbin.org/post"

	// TODO: POST operation using Post
	requestBody := strings.NewReader(`
	{
		"field1": "this is field1",
		"field2": "200"

	}`);

	response, err := http.Post(httpbin, "application/json", requestBody);
	if err != nil {
		return;
	}

	content, _ := ioutil.ReadAll(response.Body);
	defer response.Body.Close();
	fmt.Printf("Response :\n%s\n", content);

	// TODO: POST operation using PostForm
	data := url.Values{};
	data.Add("field1", "Field added via values");
	data.Add("field2", "300");
	response, err = http.PostForm(httpbin, data);
	if err != nil {
		return;
	}

	content, _ = ioutil.ReadAll(response.Body);
	defer response.Body.Close();
	fmt.Printf("Response to post form :\n%s\n", content);

}

func main() {
	// Execute a POST
	postRequestTest()
}
