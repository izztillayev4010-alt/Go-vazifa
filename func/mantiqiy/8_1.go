// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Uch xonali a soni berilgan jumlani rostlikka tekshiring: a soni polindrom va juf son
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
	y := x/100 == x%10 && x%2 == 0
	fmt.Println(y)
}
