package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// TODO: get the default temporary directory path
	tempPath := os.TempDir();
	fmt.Println("Temporary files are under :", tempPath);

	// TODO: create a temporary file use TempFile
	tempContent := []byte("This is somme dummy content");
	tempFile, err := ioutil.TempFile(tempPath, "temp_*.txt");
	if err != nil {
		panic(err);
	}

	if _, err = tempFile.Write(tempContent); err != nil {
		panic(err);
	}

	// TODO: Read and print the tempFile content
	data, _ := ioutil.ReadFile(tempFile.Name());
	fmt.Printf("%s\n", data);

	// TODO: Remove the temp file when finished
	fmt.Println("Remove temp file", tempFile.Name());
	defer os.Remove(tempFile.Name());

	// TODO: create a temporary directory with ioutil's version of TempDir
	tempDir, err := ioutil.TempDir("", "tempdir*");
	if err != nil {
		panic(err);
	}
	fmt.Println("Temp Directory :", tempDir);
	defer os.RemoveAll(tempDir);

}
