package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")
	fmt.Println()
	fmt.Println("1+1=", 1+1)
	fmt.Println("7/3=", 7.0/3.0)
	fmt.Println("7/3=", 7/3)
	fmt.Println()
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

// go run 01-value/main.go
