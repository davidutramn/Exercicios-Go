package main

import "fmt"

func main() {
	const (
		a = 1
		b
		c = 2
		d
	)
	const e = 1
	// const e = retorna1() // erro

	const (
		f = iota * 1024
		g
		h
		i
		j
	)

	fmt.Println(a, b, c, d, e)
	fmt.Println(f, g, h, i, j)
}

func retorna1() int {
	return 1
}
