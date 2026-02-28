// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: a va b sonlarni berilgan jumlani rostlikka tekshiring: a soni b soniga teng
// (Qiymat qaytaradigan func bilan)

package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("a = ")
	fmt.Scan(&a)

	fmt.Print("b = ")
	fmt.Scan(&b)

	c := hisob(a, b)

	fmt.Println(c)

}
func hisob(x int, y int) bool {
	z := x == y
	return z
}
