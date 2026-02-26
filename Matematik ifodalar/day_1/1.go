package main

import "fmt"

func main() {
	var a, b, s, p float32
	fmt.Print("a =")
	fmt.Scan(&a)
	fmt.Print("b = ")
	fmt.Scan(&b)
	s = a * b
	p = 2 * (a + b)
	fmt.Println("yuzasi: ", s, "perimetri: ", p)
}
