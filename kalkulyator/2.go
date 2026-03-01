package main

import "fmt"

func main() {
	var a int
	fmt.Print("a = ")
	fmt.Scan(&a)

	tub(a)

}
func tub(x int) {

	for i := 2; i < x; i++ {
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
