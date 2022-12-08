package main

import "fmt"

type dog struct{}

func (d dog) Cry() { fmt.Println("Muang muang") }

type cat struct{}

func (c cat) Cry() { fmt.Println("Meong meong") }

type animal interface {
	Cry()
}

// touch accepts different types
func touch(a animal) { a.Cry() }

func main() {
	dada := dog{}
	touch(dada)

	lily := cat{}
	touch(lily)
}
