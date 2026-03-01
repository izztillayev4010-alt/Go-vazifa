package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("a = ")
	fmt.Scan(&a)

	fmt.Print("b = ")
	fmt.Scan(&b)

	tub(a, b)

}
func tub(x, y int) {

	for i := 2; i >= y && i <= x; i++ {
		tubmi := true
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				tubmi = false
			}

		}
		if tubmi {
			fmt.Println(i)
		}
	}

}
