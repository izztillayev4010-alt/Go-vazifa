// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Ikki xonal a soni berilgan! Jumlani rostlikka tekshiring: a soni polindrom va toq son
// (Qiymat qaytaradigan func bilan)

package main

import "fmt"

func main() {
	var a int

	fmt.Print("a = ")
	fmt.Scan(&a)

	b := hisob(a)
	fmt.Println(b)
}
func hisob(x int) bool {
	y := x/10 == x%10 && x%2 != 0
	return y
}
