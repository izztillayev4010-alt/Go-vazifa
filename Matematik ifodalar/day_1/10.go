package main

import "fmt"

func main() {
	var a, b, c, d, e, f int
	fmt.Print("a = ")
	fmt.Scan(&a)
	b = a / 1000
	c = (a / 100) % 10
	d = (a / 10) % 10
	e = a % 10
	f = e*1000 + d*100 + c*10 + b
	fmt.Println("o'rin almashgani: ", f)
}
