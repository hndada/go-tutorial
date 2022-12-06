package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func main() {
	result1 := plus(1, 2)
	fmt.Println(result1)

	result2 := minus(1, 2)
	fmt.Println(result2)
}
