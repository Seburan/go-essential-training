package main

import (
	"fmt"
	"os"
)

func checkFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return false;
		}
	}
	return true;
}

func main() {
	filePath := "sampletext.txt";

	// Use the Stat function to get file stats
	stats, err := os.Stat(filePath);
	if err != nil {
		panic(err);
	}

	// Check if a file exists
	exists := checkFileExists(filePath);
	fmt.Printf("%s : %t\n", filePath, exists);

	// Get the file's modification time
	fmt.Println("File modification time :", stats.ModTime());
	fmt.Println("File mode:", stats.Mode());
	fMode := stats.Mode();
	if fMode.IsRegular() {
		fmt.Println("This is a regular file");
	}

	// Get the file size
	fmt.Println("File size :", stats.Size());

	// Is this a directory?
	fmt.Println("Is it a directory ? :", stats.IsDir());

}
