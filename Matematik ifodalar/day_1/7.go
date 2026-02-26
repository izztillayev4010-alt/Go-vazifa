package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Print("a = ")
	fmt.Scan(&a)
	b = a / 10
	c = (a % 10)
	d = c*10 + b
	fmt.Println("o'rin almashgani: ", d)
}
