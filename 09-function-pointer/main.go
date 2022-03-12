package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func setByValue(n int) {
	n = 5
}
func setByPointer(n *int) {
	*n = 5
}

func main() {
	result1 := plus(1, 2)
	fmt.Println(result1)

	result2 := minus(1, 2)
	fmt.Println(result2)
	fmt.Println()

	// pointer
	n := 10
	fmt.Println("n")
	fmt.Println("before:", n)
	setByValue(n)
	fmt.Println("after:", n)

	m := 10
	fmt.Println("m")
	fmt.Println("before:", m)
	setByPointer(&m)
	fmt.Println("after:", m)

}

// go run 09-function-pointer/main.go
