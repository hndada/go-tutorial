package main

import "fmt"

func main() {
	m := make(map[string]int) // beware: make
	m["a"] = 10
	m["b"] = 50
	fmt.Println(`value of key "a": `, m["a"])
	fmt.Println("length:", len(m))
	fmt.Println(m)
	fmt.Println()

	delete(m, "b")
	fmt.Println(m)

	v1 := m["a"]
	fmt.Println(v1)

	v2, ok := m["b"] // ok is optional
	fmt.Println(v2, ok)
}
