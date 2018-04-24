// assignment 3
// read the contents of a text file and print it out
// so ... cat?

package main

import (
	"fmt"
	"io"
	"os"
)

func readFile(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Unable to open file: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
	file.Close()
}

func main() {
	readFile(os.Args[1])
}
