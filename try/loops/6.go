/*
Sana: 21-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Berilgan n sonini raqamlarini teskari tartib yozishda xosil bo’lgan sonni topuvchi dastur tuzing
*/

package main

import "fmt"

func main() {
	var n int

	fmt.Print("n = ")
	fmt.Scan(&n)

	reverse := 0

	for n > 0 {
		r := n % 10

		reverse = reverse*10 + r

		n = n / 10
	}
	fmt.Println(reverse)
}
