// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: a va b sonlarni berilgan jumlani rostlikka tekshiring: a soni b sonidan katta
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
	z := x > y
	fmt.Println(z)
}
