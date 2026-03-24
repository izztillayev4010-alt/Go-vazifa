/*
sana: 19-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: a va b o’zgaruvchilar berilgan. Ularning qiymatlarni o’rini almashtiruvhi dastur tuzing! (2 xil usulda yeching)
*/

package main

import "fmt"

func main() {
	var (
		a, b int
	)

	fmt.Print("a = ")
	fmt.Scan(&a)

	fmt.Print("b = ")
	fmt.Scan(&b)

	//1-usul
	c := a
	a = b
	b = c

	fmt.Println("a = ", a, " ", "b = ", b)

	//2-usul
	a, b = b, a

	fmt.Println("a = ", a, " ", "b = ", b)
}
