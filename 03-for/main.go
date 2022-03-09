package main

import "fmt"

func main() {
	i := 0
	for i < 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println()

	for {
		fmt.Println("loop")
		break
	}
	fmt.Println()

	for j := 5; j < 8; j++ {
		fmt.Println(j)
	}
	fmt.Println()

	for n := 0; n < 7; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

// go run 03-for/main.go
