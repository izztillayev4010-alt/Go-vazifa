// Sana: 20-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: 1 dan 10 gacha bo’lgan sonlarni teskari tartibda ekranga chiqaruchi dastur tuzing

package main

import "fmt"

func main() {
	var a int

	a = 10

	for i := 10; i >= 1; i-- {
		fmt.Println("a = ", a)
		a--
	}
}
