package main

import "fmt"

func main() {
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

	p := make([]string, 0, 3)
	fmt.Println(p)
	p = append(p, "x")
	fmt.Println(p)

}

// go run tutorial/06-slice/main.go
