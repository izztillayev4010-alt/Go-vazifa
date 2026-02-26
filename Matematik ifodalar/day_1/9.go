package main

import "fmt"

func main() {
	var a, b, c, d, e, f, h, g int
	fmt.Print("a = ")
	fmt.Scan(&a)
	b = a % 10
	c = (a / 100) % 10
	d = b + c
	e = a / 1000
	f = (a / 10) % 10
	h = e + f
	g = d - h
	fmt.Println("ayirma: ", g)

}
