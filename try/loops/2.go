/*
Sana: 21-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: 1 dan 10 gacha bo’lgan sonlarni kvadratlarini ekranga chiqaruchi dastur tuzing
*/

package main

import "fmt"

func main() {
	var a []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var b int

	for i := 0; i < len(a); i++ {
		b = a[i] * a[i]
		fmt.Println(b)
	}

}
