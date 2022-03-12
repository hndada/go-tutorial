package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")
	fmt.Println()
	fmt.Println("1+1=", 1+1)
	fmt.Println("7/3=", 7.0/3.0) // float / float = float
	fmt.Println("7/3=", 7/3)     // int / int = int
	fmt.Println()
	fmt.Println(true && false) // and
	fmt.Println(true || false) // or
	fmt.Println(!true)         // not
}

// go run 01-value/main.go
