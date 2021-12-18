// Challenge : Channels
package main

import (
	"fmt"
	"net/http"
)

func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s -> %s\n", url, ctype)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	// var wg sync.WaitGroup;
	var ch = make(chan int); // create a channel

	for _, url := range urls {

		go func(url string) {
			returnType(url)
			ch <- 1; // push index number to channel
		} (url);
		//close(ch);
	}

	for i, _ := range urls {
		val := <-ch;
		fmt.Printf("%d - received %d\n", i, val);
	}
}
