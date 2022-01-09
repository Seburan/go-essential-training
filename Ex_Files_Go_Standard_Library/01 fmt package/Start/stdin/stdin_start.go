package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin);
	s, err := reader.ReadString('\n');
	fmt.Println(s);
	if err != nil {
		fmt.Println("Error : ", err);
	}
}
