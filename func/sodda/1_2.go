// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: To’gri to’rtburchakni a va b tomonlari berilgan uning yuzini va perimetrini topuvchi dastur tuzing!
//(Qiymat qaytaruvchi func bilan)

package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("a = ")
	fmt.Scan(&a)

	fmt.Print("b = ")
	fmt.Scan(&b)

	s := hisob1(a, b)
	p := hisob2(a, b)

	fmt.Println("S = ", s, " ", "P = ", p)
}

func hisob1(x int, y int) int {
	m := x * y

	return m
}
func hisob2(x int, y int) int {
	n := 2 * (x + y)

	return n
}
