package main

import (
	"fmt"
	"sync"
	"time"
)

func hard() {
	for i := 0; i < 2*1e9; i++ {
	}
	fmt.Println("hard work done")
}
func easy() {
	for i := 0; i < 10; i++ {
	}
	fmt.Println("easy work done")
}
func main() {
	go hard()
	easy()
	time.Sleep(2 * time.Second) // need to wait hard()
	fmt.Println()

	// wait groups: nicer than time.Sleep()
	var wg sync.WaitGroup
	wg.Add(1) // wg's counter: 0 -> 1
	go func() {
		defer wg.Done() // wg's counter: 1 -> 0
		hard()
	}()

	easy()
	wg.Wait() // wait until wg's counter is 0
}
