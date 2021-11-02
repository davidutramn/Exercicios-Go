package main

import "fmt"

const (
	one = iota + 1
	two
	three
)

func main() {
	var a [3]int
	b := [3]int{1, 2, 3}          // array literal
	c := [...]int{1, 2, 3}        // the size is the number of elements
	d := [5]int{1, 2}             // the rest is the zero value of the type (0 to int)
	e := [5]int{0: 1, 4: 2}       // specify a list of index
	f := [5]int{one: 1, three: 2} // can use constants to specify a index
	g := [...]int{99: 1}          // creates the array with the specified last index

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	fmt.Println(a == b)
	fmt.Println(b == c)
	// fmt.Println(b == d) // compile error: mismatched types [3]int and [5]int
	fmt.Println(d == e)
}
