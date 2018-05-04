// create a buffered channel, fill it with values and close
// test of reading from a closed channel as well as simple WaitGroup

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan int, 10)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for x := 10; x < 20; x++ {
			c <- x
		}
		close(c)
		wg.Done()
	}()

	go func() {
		// odd: if range c is used without the variable
		// every _other_ value will be output from the channel
		for n := range c {
			time.Sleep(time.Second)
			fmt.Println(n)
		}
		wg.Done()
	}()

	wg.Wait()
}
