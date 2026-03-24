/*
sana: 19-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: To’gri to’rtburchakni a va b tomonlari berilgan uning yuzini va perimetrini topuvchi dastur tuzing!
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

	s := a * b

	p := 2 * (a + b)

	fmt.Println("S = ", s, " ", "P = ", p)
}
