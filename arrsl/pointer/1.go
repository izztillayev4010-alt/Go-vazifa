// Sana: 13-03-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: To’gri to’rtburchakni a va b tomonlari berilgan uning yuzini va perimetrini topuvchi dastur tuzing!
//(pointer bilan)

package main

import "fmt"

func main() {
	var a, b int
	var s, p int

	fmt.Print("a = ")
	fmt.Scan(&a)

	fmt.Print("b = ")
	fmt.Scan(&b)

	hisob(&a, &b, &s, &p)

	fmt.Println("Yuzi: ", s, " ", "Perimetri: ", p)

}
func hisob(p1, p2, r, t *int) {
	*r = (*p1) * (*p2)
	*t = 2 * (*p1 + *p2)
}
