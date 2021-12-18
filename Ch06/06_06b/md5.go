/*
Calculate MD5 signature of several files concurrently. your input is a
file where every line has an expected signature and a file name.

you need to parse the file, and then for every file calculate the actual
md5 signature of the file. it it doesn't match the expeted value,
print an error message
*/
package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const logpath string = "nasa-logs/";

// Parse the signature file, return a map of path->signature
func parseSignatureFile(path string) (map[string]string, error) {
	// read md5sum.txt content
	file, err := os.Open(path);
	if err != nil {
		return nil, err;
	}
	// close file descriptor
	defer file.Close();

	// create a map of path->signature
	var signatures = make(map[string]string);

	// parse md5sum.txt content
	scanner := bufio.NewScanner(file);
	for lineIndex :=0; scanner.Scan(); lineIndex++ {

		tokens := strings.Fields(scanner.Text())
		if len(tokens) != 2 {
			return nil, fmt.Errorf("file format expects md5sum and filename only. review file <%s> Line <%d>", path, lineIndex);
		}

		// save filename and its signature
		signatures[logpath + tokens[1]] = tokens[0];

	}

	if err := scanner.Err(); err != nil {
		return nil, err;
	}

	return signatures, nil;
}

// return the md5sum of the file
func fileMD5(path string) (string, error) {
	// read logfile and compute md5hash
	file, err := os.Open(path);
	if err != nil {
		return "", err;
	}
	// close file descriptor
	defer file.Close();

	// compute md5sum from file content
	h := md5.New();
	if _, err := io.Copy(h, file); err != nil {
		return "", err;
	}

	checksum := fmt.Sprintf("%x", h.Sum(nil));

	return checksum, nil;
}

// check file integrity against given expected signature and send result back
// to channel ch
func integrityCheck(path string, expectedSignature string, ch chan string) {

	// calculte md5sum signature of the file
	signature, err := fileMD5(path);
	if err != nil {
		ch <- fmt.Sprintf("%s", err);
	}

	// check signature against expected signature
	if signature != expectedSignature {
		ch <- fmt.Sprintf("%s :  %s (expected value : %s)", path, signature, expectedSignature);
	} else {
		ch <- fmt.Sprintf("%s :  %s (OK)", path, signature);
	}

}

func main() {

	// setup channel
	var ch = make(chan string);

	// parse the signature file and store them in map[path]->signature
	signatures, err := parseSignatureFile(logpath + "md5sum.txt");
	if err != nil {
		log.Fatalf("Error : cannot parse the signature file. %s", err)
	}

	// run integrity check for each file in a concurrent routine
	for path, signature := range signatures {
		go integrityCheck(path, signature, ch);
	}

	// show checksum results for each file in the signature file
	for range signatures {
		fmt.Printf("%s\n",<-ch);
	}

}
