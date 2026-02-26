package main

import "fmt"

func main() {
	var c float32
	var a, b, d int
	fmt.Print("a = ")
	fmt.Scan(&a)
	b = a / 10
	c = float32(a % 10)
	d = b + int(c)
	fmt.Println("Yig'indi: ", d)
}
