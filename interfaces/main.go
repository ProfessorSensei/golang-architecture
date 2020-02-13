package main

import "fmt"

type person struct {
	first string
}

func main() {
	p1 := person{
		first: "Professor Sensei",
	}

	fmt.Printf("%T\n", p1)
}
