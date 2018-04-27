// code-along
// experimenting with concurrency
// preferred : channels!
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64
var mutex sync.Mutex

func main() {
	rand.Seed(time.Now().UnixNano())
	wg.Add(2)
	// go foo()
	// go bar()
	go incrementor("Foo")
	go incrementor("Bar")
	wg.Wait()
	fmt.Println("Final counter:", counter)
}

func foo() {
	for i := 0; i < 50; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 50; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(4 * time.Millisecond)
	}
	// forget a wg.Done() ? Program crashes with a fatal deadlock
	wg.Done()
}

// intentionally broken func + constant to see both race condition
// and it's detection by `go run -race`
// go run -race shows race condition even when using mutex if x is still in use
func incrementor(title string) {
	for i := 0; i < 20; i++ {
		// mutex.Lock()
		// x := counter
		// x++
		time.Sleep(time.Duration(rand.Int63n(5)) * time.Millisecond)
		// counter = x
		// counter++

		// other way : sync/atomic
		atomic.AddInt64(&counter, 1)

		// all uses of the counter variable must be in the mutex
		// otherwise go run -race reports a data race
		// despite this being a print statement
		// if mutex.Unlock() is above it, still a data race.
		// or, if atomic is used but this shared variable is printed ... race
		fmt.Println(title, i, "counter: ", counter)
		// mutex.Unlock()
	}
	wg.Done()
}
