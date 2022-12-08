package main

import "fmt"

func main() {
	if 2 > 3 {
		fmt.Println("2 is greater than 3")
	} else if 2 == 3 {
		fmt.Println("2 is equal to 3")
	} else {
		fmt.Println("2 is less than 3")
	}
	fmt.Println()

	if i := 6; i%4 == 0 {
		fmt.Println("6 is divided by 4")
	}
	fmt.Println()

	// switch
	n := 2
	switch n {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3, 4:
		fmt.Println("three or four")
	default:
		fmt.Println("over")
	}

	switch {
	case n%2 != 0:
		fmt.Println("odd")
	default:
		fmt.Println("even")
	}
}
