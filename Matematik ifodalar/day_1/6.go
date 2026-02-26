package main

import "fmt"

func main() {
	var c float32
	var a, d, e, b int
	fmt.Print("a = ")
	fmt.Scan(&a)
	e = a / 100
	b = (a / 10) % 10
	c = float32(a % 10)
	d = b + int(c) + e
	fmt.Println("Yig'indi: ", d)
}
