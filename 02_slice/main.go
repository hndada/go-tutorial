package main

import "fmt"

func main() {
	// array: a sequence of variables
	var a [5]int
	fmt.Println(a)

	a[4] = 100
	fmt.Println(a)
	fmt.Println("len:", len(a))

	a = [5]int{0, 2, 4, 6, 8}

	var two [2][3]int
	fmt.Println(two)

	// slice: a powerful array
	s := make([]int, 3)
	s[2] = 20
	fmt.Println(s)

	s = append(s, 30)
	s = append(s, 40, 50)
	fmt.Println(s)

	l := s[2:5]
	fmt.Println("s 2~5:", l)
	l = s[:5]
	fmt.Println("s ~5:", l)
	l = s[2:]
	fmt.Println("s 2~:", l)

	t := []string{"a", "b", "c"}
	fmt.Println(t)
}
