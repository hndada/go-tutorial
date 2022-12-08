package main

import "fmt"

func main() {
	m := make(map[string]int) // beware: make
	m["a"] = 10
	m["b"] = 50
	fmt.Println(`value of key "a": `, m["a"])
	fmt.Println("length:", len(m))
	fmt.Println(m)

	delete(m, "b")
	fmt.Println(m)

	v, ok := m["b"]
	fmt.Println(v, ok)
}
