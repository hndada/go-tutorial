package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func setValue(n int) {
	n = 5
}
func setPointer(n *int) {
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
	setValue(n)
	fmt.Println("after:", n)

	m := 10
	fmt.Println("m")
	fmt.Println("before:", m)
	setPointer(&m)
	fmt.Println("after:", m)

}

// go run tutorial/09-function-pointer/main.go
