// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Berilgan ikki xonali sonni raqamlari yig’indisini toping!
//(Qiymat qaytarmaydigan func bilan)

package main

import "fmt"

func main() {
	var a int

	fmt.Print("a = ")
	fmt.Scan(&a)

	hisob(a)
}
func hisob(x int) {
	y := x / 10
	z := x % 10
	fmt.Println("yig'indi: ", y+z)
}
