package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("a = ")
	fmt.Scan(&a)
	fmt.Print("b = ")
	fmt.Scan(&b)
	c = a
	a = b
	b = c
	fmt.Println("a =", a, " ", "b = ", b)

}
