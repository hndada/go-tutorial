package main

import "fmt"

func main() {
	nums := []int{1, 3, 5}
	sum := 0
	for _, n := range nums {
		sum += n
	}
	fmt.Println("sum:", sum)

	for i, n := range nums {
		fmt.Println("index:", i, "number:", n)
	}

	fruits := map[string]string{
		"a": "apple", "b": "banana"}
	for k, v := range fruits {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k := range fruits {
		fmt.Println("key:", k)
	}
}

// go run 08-range/main.go
