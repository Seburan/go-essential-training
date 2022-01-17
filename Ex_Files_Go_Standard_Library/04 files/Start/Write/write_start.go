package main

import (
	"fmt"
	"os"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func checkFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func writeFileData(filePath string) {
	// TODO: create a new file
	f, err := os.Create(filePath);
	handleErr(err);

	// TODO: it is idiomatic in Go to defer the close after you open the file
	defer f.Close();

	// TODO: get the Name of the file
	fmt.Println("The file name is", f.Name());

	// TODO: write some string data to the file
	f.WriteString("This is some text\n");

	// TODO: write some individual bytes to the file
	data2 := []byte{'a','b', 'c', '\n'};
	f.Write(data2);

	// TODO: the Sync function forces the data to be written out
	f.Sync();
}

func appendFileData(fname string, data string) {
	// TODO: Use the lower-level OpenFile function to open the file in append mode
	f, err := os.OpenFile(fname,os.O_APPEND|os.O_WRONLY, 0644);
	handleErr(err);
	defer f.Close();

	_, err = f.WriteString(data);
}

func main() {

	filePath := "datafile.txt";

	// TODO: Simple example: dump some data to a file
	// data1 := []byte("This is some text data\n");
	// ioutil.WriteFile(filePath, data1, 0644);

	// TODO: More complex example: Granular writing of data
	if !checkFileExists(filePath) {
		writeFileData(filePath);
	} else {
		os.Truncate(filePath, 10);
		fmt.Println("Trimmed the file data");
	}

	// TODO: append data to a file
	appendFileData(filePath, "This data was appended");

}
