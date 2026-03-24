/*
Sana: 20-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: a, b va c sonlari berilgan. Bu sonlarning o’rachasini topuvchi dastur tuzing!
*/

package main

import "fmt"

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
	r := 0
	if a > b {
		res = a
		r = b
	} else if b > a {
		res = b
		r = a
	} else {
		res = a

	}
	if res > c {
		if c > r {
			fmt.Println("(a,b,c)= ", c)
		} else if r > c {
			fmt.Println("(a,b,c)= ", r)
		}
	} else if res < c {
		fmt.Println("(a,b,c)= ", res)
	} else {
		fmt.Println("(a,b,c)= ", res)
	}
}
