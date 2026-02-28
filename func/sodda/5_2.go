// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Berilgan ikki xonali sonni raqamlari yig’indisini toping!
//(Qiymat qaytaradigan func bilan)

package main

import "fmt"

func main() {
	var a int

	fmt.Print("a = ")
	fmt.Scan(&a)

	c := hisob(a)

	fmt.Println("yig'indi: ", c)
}
func hisob(x int) int {
	y := x / 10
	z := x % 10
	b := y + z
	return b
}
