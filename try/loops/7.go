/*
Sana: 21-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Berilgan n sonini poindrom yoki polindrom emasligini aniqlovchi dastur tuzing!
*/

package main

import "fmt"

func main() {
	var n int

	fmt.Print("n = ")
	fmt.Scan(&n)

	reverse := 0
	org := n

	for n > 0 {
		r := n % 10

		reverse = reverse*10 + r

		n = n / 10
	}
	if reverse == org {
		fmt.Println("n polindrom! ", org)
	} else {
		fmt.Println("polindrom emas! ")
	}
}
