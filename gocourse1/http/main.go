package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com/")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// don't want to print Resp by itself
	// the manual way with specific byte slice
	// initialized large so that Read() doesn't quit
	// when slice is full
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// short form without using Reader interface directly
	// io.Copy -> 32K buffer
	// io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

// example of a Writer
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
