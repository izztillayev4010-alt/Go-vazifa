// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Berilgan a soni musbatligini aniqlovchi dastur tuzing
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
	y := x > 0
	return y
}
