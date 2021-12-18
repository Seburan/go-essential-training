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

func integrityCheck(inputChecksum string, filename string, ch chan string) {

	// read logfile and compute md5hash
	file, err := os.Open(logpath + filename);
	if err != nil {
		log.Fatal(err);
	}
	defer file.Close();

	// compute md5sum from file content
	h := md5.New();
	if _, err := io.Copy(h, file); err != nil {
		log.Fatal(err);
	}

	checksum := fmt.Sprintf("%x", h.Sum(nil));

	if checksum != inputChecksum {
		ch <- fmt.Sprintf("%s :  %s (expected value : %s)", filename, checksum, inputChecksum);
	} else {
		ch <- fmt.Sprintf("%s :  %s (OK)", filename, checksum);
	}


}

func main() {

	// setup channel
	var ch = make(chan string);
	var fileCount int;

	// read md5sum.txt content
	file, err := os.Open(logpath + "md5sum.txt");
	if err != nil {
		log.Fatal(err);
	}

	// parse md5sum.txt content
	scanner := bufio.NewScanner(file);
	for scanner.Scan(){
		// fmt.Println(scanner.Text())
		tokens := strings.Fields(scanner.Text())
		if len(tokens) != 2 {
			fmt.Fprintln(os.Stderr, "File format expects md5sum and filename only", err)
		}
		md5 := tokens[0];
		filename := tokens[1];
		fileCount++;
		// fmt.Printf("filename : %s - md5 : %s\n", filename, md5);
		go integrityCheck(md5, filename, ch);
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading file :", err)
	}

	// close file descriptor
	if err := file.Close(); err != nil {
		log.Fatal(err);
	}

	// show checksum results
	for i := 0; i < fileCount; i++ {
		fmt.Printf("%s\n",<-ch);
	}

}
