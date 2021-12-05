// Write a truct called Capper that has a field to another io.Writer and
// transforms everything written to uppercase. Capper should implement
// io.Writer
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type Capper struct {
	wtr io.Writer;
}

func (c *Capper) Write(p []byte) (n int, err error) {
	// overload the io.Writer method and upper case the written bytes
	n, err = c.wtr.Write(bytes.ToUpper(p));

	return n, err;
}

func main() {
	fmt.Println("Challenge : Interfaces")

	// Create new Capper using os.Stdout as its underlying writer
	w := &Capper{os.Stdout};
	n, err := fmt.Fprintln(w, "Hello World !");

	// The n and err return values from Fprintln are
	// those returned by the underlying io.Writer.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintln: %v\n", err);
	}
	fmt.Println(n, "bytes written.")

}
