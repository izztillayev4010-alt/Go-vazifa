// Sana: 13-03-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: a va b o’zgaruvchilar berilgan. Ularning qiymatlarni o’rini almashtiruvhi dastur tuzing! (2 xil usulda yeching)
//(pointer bilan)

package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("a = ")
	fmt.Scan(&a)

	fmt.Print("b = ")
	fmt.Scan(&b)

	hisob(&a, &b)

	fmt.Println("a = ", a, " ", "b = ", b)

}
func hisob(m, n *int) {
	*m, *n = *n, *m

}
