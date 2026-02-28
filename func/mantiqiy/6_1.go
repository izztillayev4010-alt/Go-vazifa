// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: a va b sonlari berilgan! Bu sonlarning yig’indisi musbatligini va toqligini aniqlovchi dastur tuzing!
// (Qiymat qaytarmaydigan func bilan)

package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("a = ")
	fmt.Scan(&a)

	fmt.Print("b = ")
	fmt.Scan(&b)

	hisob(a, b)

}
func hisob(x int, y int) {
	z := (x+y) > 0 && (x+y)%2 != 0
	fmt.Println(z)
}
