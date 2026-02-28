// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: a va b o’zgaruvchilar berilgan. Ularning qiymatlarni o’rini almashtiruvhi dastur tuzing! (2 xil usulda yeching)
//(Qiymat qaytaradigan func bilan)

package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("a = ")
	fmt.Scan(&a)

	fmt.Print("b = ")
	fmt.Scan(&b)
	m, n := almashtirish(a, b)
	fmt.Println("Almashgani: ", m, n)

}
func almashtirish(x int, y int) (int, int) {
	c := x
	x = y
	y = c
	return x, y
}
