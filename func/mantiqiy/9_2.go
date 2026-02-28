// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: A soni berilgan! Jumlani rostlikka tekshiring! A soni 3 xonali va toq son
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
	y := x > 99 && x < 1000 && x%2 != 0
	return y
}
