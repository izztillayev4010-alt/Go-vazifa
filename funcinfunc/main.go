package main

import "fmt"

func main() {
	fn1 := func(x, y int) int {
		return x + y
	}
	fn2 := func(x, y int) int {
		return x - y
	}

	resfn := fn(fn1, fn2)

	res := resfn(0, 0)

	fmt.Println(res)

}

func fn(fn1 func(x, y int) int, fn2 func(x, y int) int) func(x, y int) int {
	var a, b int

	fmt.Print("a = ")
	fmt.Scan(&a)

	fmt.Print("b = ")
	fmt.Scan(&b)

	r1 := fn1(a, b)
	r2 := fn2(a, b)

	if (r1+r2)%2 == 0 {
		fn3 := func(x, y int) int {
			return r1 + r2
		}
		return fn3
	} else {
		fn4 := func(x, y int) int {
			return r1 - r2
		}
		return fn4
	}
}
