package main

import "fmt"

func main() {
	i := 0
	for i < 3 {
		fmt.Println(i)
		i = i + 1 // add 1 to i
	}
	fmt.Println()

	for {
		fmt.Println("loop")
		break
	}
	fmt.Println()

	for j := 5; j < 8; j++ { // j = j + 1
		fmt.Println(j)
	}
	fmt.Println()

	for n := 0; n < 7; n++ {
		if n < 3 {
			continue
		}
		fmt.Println(n)
	}
}

// go run 03-for/main.go
