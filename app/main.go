package main

import (
	"fmt"

	"mod.go/myMath"
	"mod.go/myStr"
)

func main() {
	a := myStr.Count("Banana", "a")
	b := myMath.Abs(-43)
	fmt.Println(a, b)
	c := myMath.Pow10(2)
	fmt.Println(c)
}
