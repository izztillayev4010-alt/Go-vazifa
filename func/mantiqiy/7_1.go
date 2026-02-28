// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Ikki xonal a soni berilgan! Jumlani rostlikka tekshiring: a soni polindrom va toq son
// (Qiymat qaytarmaydigan func bilan)

package main

import "fmt"

func main() {
	var a int

	fmt.Print("a = ")
	fmt.Scan(&a)

	hisob(a)
}
func hisob(x int) {
	y := x/10 == x%10 && x%2 != 0
	fmt.Println(y)
}
