// Sana: 20-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: 1 dan 10 gacha bo’lgan sonlarni kvadratlarini ekranga chiqaruchi dastur tuzing

package main

import "fmt"

func main() {
	var a int

	a = 1

	for i := 0; i < 10; i++ {
		c := a * a
		fmt.Println("a = ", c)
		a++
	}
}
