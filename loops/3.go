// Sana: 20-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: 1 dan 10 gacha bo’lgan sonlarni yig’indisni ekranga chiqaruchi dastur tuzing

package main

import "fmt"

func main() {
	var a int

	a = 1
	sum := 0

	for i := 0; i < 10; i++ {
		sum = sum + a

		a++
	}
	fmt.Println("a = ", sum)
}
