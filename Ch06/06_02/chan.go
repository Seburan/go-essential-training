// channels
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	// This will block
	/*
		<-ch
		fmt.Println("Here")
	*/

	go func() {
		// Send number to the channel
		ch <- 353
	} ();

		// Receive from the channel
		val := <-ch;
		fmt.Printf("got %d\n", val);

		fmt.Println("--------");

		// Send multiple numbers to the channel
		go func() {
			for i := 0; i < 3; i++ {
				fmt.Printf("sending %d\n", i);
				ch <- i;
				time.Sleep(time.Second);
			}
			}();

			for i := 0; i < 3; i++ {
				val := <-ch;
				fmt.Printf("received %d\n", val);
			}

		fmt.Println("--------");

		// Close to signal we are done
		go func() {
			for i := 0; i < 3; i++ {
				fmt.Printf("sending %d\n", i);
				ch <- i;
				time.Sleep(time.Second);
			}
			close(ch);
		}();

		for i := range ch {
			fmt.Printf("received %d\n", i);
		}

}
