package main

import "fmt"

func main() {
	var a, b, c, d, e int
	fmt.Print("a = ")
	fmt.Scan(&a)
	b = a / 100
	c = (a / 10) % 10
	d = a % 10
	e = c*100 + b*10 + d
	fmt.Println("o'rin almashgani: ", e)
}
