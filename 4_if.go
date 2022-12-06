package main

import "fmt"

func main() {
	if 2 > 3 {
		fmt.Println("2 is larger than 3")
	} else {
		fmt.Println("2 is smaller than 3")
	}

	if 1+1 == 2 {
		fmt.Println("of course 1+1 is 2")
	}
	fmt.Println()

	fmt.Print("6 is ")
	if i := 6; i%5 == 0 {
		fmt.Println("5 multiple")
	} else if i%4 == 0 {
		fmt.Println("4 multiple")
	} else if i%3 == 0 {
		fmt.Println("3 multiple")
	} else if i%2 == 0 {
		fmt.Println("2 multiple")
	} else {
		fmt.Println("a prime")
	}
	fmt.Println()

	// switch
	n := 2
	switch n {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("over")
	}

	switch n {
	case 1, 3:
		fmt.Println("odd")
	case 2, 4:
		fmt.Println("even")
	}

	switch {
	case n%2 != 0:
		fmt.Println("odd")
	default:
		fmt.Println("even")
	}
}
