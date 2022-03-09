package main

import "fmt"

func main() {
	names := []string{"value", "variable", "for", "if", "array", "slice", "map", "range", "function-pointer", "struct-method"}
	for i, name := range names {
		fmt.Printf("go run tutorial/%02d-%s/main.go\n", i+1, name)
	}
}
