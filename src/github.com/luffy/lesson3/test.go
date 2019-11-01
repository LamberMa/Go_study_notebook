package main

import (
	"fmt"
)

func main(){
	const(
		a = iota
		b
		c
		d = 8
		e
		f =iota
		g
	)

	fmt.Printf("a=%d, b=%d, c=%d, d=%d, e=%d, f=%d, g=%d\n", a, b, c, d, e, f, g)
}