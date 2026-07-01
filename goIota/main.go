package main

import "fmt"

type MyStatus string

const (
	pending   MyStatus = "Pending"
	rejected  MyStatus = "Rejected"
	Submitted MyStatus = "Submitted"
)

func main() {
	const (
		a = iota // always starts with 0
		b
		c = iota + 1 // here c hold value 3 because iota is already having value 2 so 2 + 1
		d
		e
		f
	)

	fmt.Println(a, b, c, d, e, f)
	var s MyStatus = pending

	fmt.Println(s)

}
