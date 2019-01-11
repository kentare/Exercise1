// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
	"time"
)

var i = 0

func incrementing() {
	//TODO: increment i 1000000 times
	for a := 0; a < 100000; a++ {
		i++
	}
}

func decrementing() {
	//TODO: decrement i 1000000 times
	for b := 0; b < 100000; b++ {
		i--
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // I guess this is a hint to what GOMAXPROCS does...
	// Try doing the exercise both with and without it!
	// Kent: GOMAXPROCS sets the maximum number of CPUs that can be executing simultaneously and returns the previous setting. If n < 1, it does not change the current setting.
	// The number of logical CPUs on the local machine can be queried with the NumCPU.

	// TODO: Spawn both functions as goroutines
	go incrementing()
	go decrementing()

	// We have no way to wait for the completion of a goroutine (without additional syncronization of some sort)
	// We'll come back to using channels in Exercise 2. For now: Sleep.
	time.Sleep(100 * time.Millisecond)
	Println("The magic number is:", i)
}
