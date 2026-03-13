// Sana: 13-03-2026
// Muallif: Izzatillayev Komronbek
// Maqsad:Doiraning radiusi berilgan r, uning aylanasi uzunligi va yuzini hisblovchi dastur tuzing
//(pointer bilan)

package main

import "fmt"

func main() {
	var r, l, s float64

	const p = 3.14

	fmt.Print("r = ")
	fmt.Scan(&r)

	b, c := hisob(&r, &l, &s)
	b = b * p
	c = c * p

	fmt.Println("L = ", b, " ", "S = ", c)

}
func hisob(m, n, t *float64) (float64, float64) {

	*n = 2 * (*m)
	*t = *m * (*m)
	return *n, *m
}
