package main

import "fmt"

func main() {
	for j := 5; j < 8; j++ { // j = j + 1
		fmt.Println(j)
	}
	fmt.Println()

	i := 0
	for i < 3 { // condition-only
		fmt.Println(i)
		i++
	}
	fmt.Println()

	for {
		fmt.Println("loop")
		break
	}
	fmt.Println()

	for n := 0; n < 7; n++ {
		if n < 3 {
			continue
		}
		fmt.Println(n)
	}

	// range
	nums := []int{1, 3, 5}
	for i, n := range nums {
		fmt.Println("index:", i, "number:", n)
	}

	sum := 0
	for _, n := range nums {
		sum += n
	}
	fmt.Println("sum:", sum)

	fruits := map[string]string{
		"a": "apple",
		"b": "banana",
	}
	for k, v := range fruits {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k := range fruits {
		fmt.Println("key:", k)
	}
}
