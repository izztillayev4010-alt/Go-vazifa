/*
Sana: 20-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: a, b va c sonlari berilgan. Bu sonlarning eng kichigini topuvchi dastur tuzing!
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
	if a > b {
		res = b
	} else if b > a {
		res = a
	} else {
		res = a
	}
	if res > c {
		fmt.Println("(a,b,c)= ", c)
	} else if res < c {
		fmt.Println("(a,b,c)= ", res)
	} else {
		fmt.Println("(a,b,c)= ", res)
	}
}
