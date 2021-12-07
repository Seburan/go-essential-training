/*
Write a killServer(pidFile string) error function that reads a process
identifier from pidFile, converts it to an integer, and prints "killing <pid>"
(instead of using os.Kill)
*/

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

// killServer reads a process identifier from pidFile, converts it to an
// integer, and prints "killing <pid>" (instead of using os.Kill)
func killServer(pidFile string) error {
		// Read the .pid file
		data, err := os.ReadFile(pidFile)
		if err != nil {
			return errors.Wrap(err, "cannot open the pid file");
		}

		if err := os.Remove(pidFile); err != nil {
			log.Printf("Warning : cannot remove pid file - %s\n", err);
		}

		dataBuffer := bytes.NewBuffer(bytes.TrimSpace(data));
		pidInt, err := strconv.Atoi(dataBuffer.String());
		if err != nil {
			return errors.Wrap(err, "cannot convert .pid file content to integer");
		}

		fmt.Printf("killing <%d>\n", pidInt);

		return nil;
}

func main() {
	err := killServer("dodo.pid");

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error : %s\n", err);
		os.Exit(1);
	}
}
