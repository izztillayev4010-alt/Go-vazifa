// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: a va b o’zgaruvchilar berilgan. Ularning qiymatlarni o’rini almashtiruvhi dastur tuzing! (2 xil usulda yeching)
//(Qiymat qaytarmaydigan func bilan)

package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("a = ")
	fmt.Scan(&a)

	fmt.Print("b = ")
	fmt.Scan(&b)
	almashtirish(a, b)

}
func almashtirish(x int, y int) {
	c := x
	x = y
	y = c
	fmt.Println("Almashgani: ", x, " ", y)
}
