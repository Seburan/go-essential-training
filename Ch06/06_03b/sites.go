// Challenge : Channels
package main

import (
	"fmt"
	"net/http"
)

func returnType(url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("%s -> error: %s\n", url, err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	ch <- fmt.Sprintf("%s -> %s\n", url, ctype)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	var ch = make(chan string); // create a channel of string

	for _, url := range urls {

		go returnType(url, ch);
		//close(ch);
	}

	for range urls {
		val := <-ch;
		fmt.Printf("%s", val);
	}
}
