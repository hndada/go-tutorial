package main

import "fmt"

type student struct {
	id   int
	name string
}

func (s student) hello() {
	fmt.Println("Hello, I'm", s.name)
}
func (s *student) setID(n int) { s.id = n }

func main() {
	s1 := student{
		id:   133,
		name: "Alice",
	}

	s2 := student{}
	s2.name = "Bob"
	fmt.Println("s1:", s1, "s2:", s2)

	// method: struct's function
	s1.hello()
	s2.hello()

	s2.setID(155)
	fmt.Println("s2:", s2)
}
