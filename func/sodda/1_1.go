// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: To’gri to’rtburchakni a va b tomonlari berilgan uning yuzini va perimetrini topuvchi dastur tuzing!
//(Qiymat qaytarmaydigan func bilan)

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
	m := x * y
	n := 2 * (x + y)
	fmt.Println("S = ", m, " ", "P = ", n)

}
