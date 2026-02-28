// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: 1 dan 10 gacha bo’lgan sonlarni teskari tartibda ekranga chiqaruchi dastur tuzing

package main

import "fmt"

func main() {
	var a int

	a = 10
	hisob(a)

}
func hisob(x int) {
	for i := 10; i >= 1; i-- {
		fmt.Println("a = ", x)
		x--
	}
}
