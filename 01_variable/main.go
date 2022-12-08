package main

import "fmt"

func main() {
	// value
	fmt.Println("go" + "lang")
	fmt.Println()

	fmt.Println("1+1=", 1+1)
	fmt.Println("7/3=", 7.0/3.0) // float / float = float
	fmt.Println("7/3=", 7/3)     // int / int = int
	fmt.Println()

	fmt.Println(true && false) // and
	fmt.Println(true || false) // or
	fmt.Println(!true)         // not
	fmt.Println()

	// variable
	var i int
	fmt.Println(i)

	var a string = "full"
	var b = "short"
	c := "shortest"
	fmt.Println(a, b, c)
}
