// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: A soni berilgan! Jumlani rostlikka tekshiring! A soni 3 xonali va toq son
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
	y := x > 99 && x < 1000 && x%2 != 0
	fmt.Println(y)
}
