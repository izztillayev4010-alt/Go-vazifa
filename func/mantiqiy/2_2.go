// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Berilgan a soni juftligini aniqlovchi dastur tuzing
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
	y := x%2 == 0
	return y
}
