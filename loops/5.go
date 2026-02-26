// Sana: 20-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: 1 dan 10 gacha bo’lgan sonlar ichidan juftlarini ekranga chiqaruchi dastur tuzing!

package main

import "fmt"

func main() {
	var a int

	a = 1

	for i := 0; i < 10; i++ {
		if a%2 == 0 {
			fmt.Print("a = ", a, " ")
		}
		a++
	}
}
