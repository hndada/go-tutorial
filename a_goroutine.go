package main

import (
	"fmt"
	"time"
)

func hard() {
	for i := 0; i < 5*1e9; i++ {
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
	time.Sleep(2 * time.Second)
}
