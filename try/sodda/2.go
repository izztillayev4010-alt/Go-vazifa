/*
sana: 19-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Doiraning radiusi berilgan r, uning aylanasi uzunligi va yuzini hisblovchi dastur tuzing
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64

	fmt.Print("r = ")
	fmt.Scan(&r)

	l := 2 * r * math.Pi

	s := math.Pow(r, 2) * math.Pi

	fmt.Println("S = ", s, " ", "L = ", l)
}
