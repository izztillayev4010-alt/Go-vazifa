/*
Sana: 20-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: a, b va c sonlari berilgan. Bu sonlarning eng kattasini topuvchi dastur tuzing!
*/

package main

import (
	"fmt"
)

func main() {
	var (
		a, b, c int
	)

	fmt.Print("a = ")
	fmt.Scan(&a)

	fmt.Print("b = ")
	fmt.Scan(&b)

	fmt.Print("c = ")
	fmt.Scan(&c)

	res := 0
	if a > b {
		res = a
	} else if b > a {
		res = b
	} else {
		res = a
	}
	if res > c {
		fmt.Println("(a,b,c)= ", res)
	} else if res < c {
		fmt.Println("(a,b,c)= ", c)
	} else {
		fmt.Println("(a,b,c)= ", res)
	}

}
