package main

import (
	"fmt"
)

func main() {
	var s, p, r, l float32
	fmt.Print("r = ")
	fmt.Scan(&r)
	p = 3.14
	s = r * r * p
	l = 2 * r * p
	fmt.Println(s, l)

}
