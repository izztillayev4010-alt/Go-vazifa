/*
Sana: 21-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Berilgan n sonini tub yoki tub emasligini aniqlovchi dastur tuzing!
*/

package main

import "fmt"

func main() {
	var n int

	fmt.Print("n = ")
	fmt.Scan(&n)

	b := true

	for i := 2; i < n; i++ {
		if n%i == 0 {
			b = false
		}
	}
	if b {
		fmt.Println("tub")
	} else {
		fmt.Println("tub emas")
	}
}
