package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	// fmt.Println("MaxProcs: ", runtime.GOMAXPROCS(0))
	// fmt.Println("CPUs found: ", runtime.NumCPU())
	links := []string{
		"http://google.com/",
		"http://stackoverflow.com/",
		"http://golang.org/",
		"http://amazon.com/",
	}

	c := make(chan string)

	for _, url := range links {
		go checkLink(url, c)
	}

	// loop through all pending goroutines
	// listening for # of messages equal to strings
	// for i := 0; i < len(links); i++ {
	// with repeated goroutines -- infinite loop
	// alternate loop syntax
	// instead of for { go checkLink(<-c, c) }
	for l := range c {
		go func(link string) {
			time.Sleep(3 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
	} else {
		fmt.Println(link, ":", resp.Status)
	}
	// pause before asking channel to launch goroutine
	// time.Sleep(3 * time.Second)
	c <- link
	return
}
