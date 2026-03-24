/*
Sana: 21-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: 1 dan 10 gacha bo’lgan sonlarni teskari tartibda ekranga chiqaruchi dastur tuzing
*/

package main

import "fmt"

func main() {
	var a []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := len(a) - 1; i >= 0; i-- {
		fmt.Println(a[i])
	}

}
