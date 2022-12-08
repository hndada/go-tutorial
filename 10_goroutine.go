package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go fmt.Println(i) // call function in parallel
	}
	time.Sleep(time.Second) // better: use WaitGroup
}
