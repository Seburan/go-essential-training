package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// TODO: create a directory
	os.Mkdir("newdir", os.ModePerm);

	// // TODO: Remove will remove an item
	// defer os.Remove("newdir");

	// TODO: create a deep directory
	os.MkdirAll("path/to/new/dir", os.ModePerm);

	// // TODO: RemoveAll will remove an item and all children
	// defer os.RemoveAll("path");

	// TODO: get the current working directory
	dir, _ := os.Getwd();
	fmt.Println("current dir is", dir);

	// TODO: get the directory of the current process
	execdir, _ := os.Executable();
	fmt.Println("exec dir is", execdir);

	// TODO: read the contents of a directory
	contents, _ := ioutil.ReadDir(".");
	for _, fi := range contents {
		fmt.Println(fi.Name(), fi.IsDir());
	}

}
