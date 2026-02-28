// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Berilgan a soni musbatligini aniqlovchi dastur tuzing
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
	y := x > 0
	fmt.Println(y)
}
