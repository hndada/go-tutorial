package main

import "fmt"

type dog struct{}
type cat struct{}

func (d dog) Cry() { fmt.Println("Muang muang") }
func (c cat) Cry() { fmt.Println("Meong meong") }

type animal interface {
	Cry()
}

// touch accepts different types
func touch(a animal) { a.Cry() }

func main() {
	dada := dog{}
	lily := cat{}
	touch(dada)
	touch(lily)
}
