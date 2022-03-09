package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 10
	m["b"] = 50
	fmt.Println(m)
	fmt.Println("len:", len(m))

	delete(m, "b")
	fmt.Println(m)

	v, ok := m["b"]
	fmt.Println(v, ok)
}

// go run tutorial/07-map/main.go
