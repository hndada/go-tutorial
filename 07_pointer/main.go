package main

import "fmt"

// Go always passes by value. Both v are new.
func setByValue(v int)    { v = 5 }
func setByPointer(v *int) { *v = 5 }

func main() {
	a := 10
	fmt.Println("a before: ", a)
	setByValue(a)
	fmt.Println("a after: ", a)

	b := 10
	fmt.Println("b before: ", b)
	setByPointer(&b)
	fmt.Println("b after: ", b)

	fmt.Println("pointer value looks like: ", &b)
}
