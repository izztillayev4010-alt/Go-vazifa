// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Doiraning radiusi berilgan r, uning aylanasi uzunligi va yuzini hisblovchi dastur tuzing
//(Qiymat qaytarmaydigan func bilan)

package main

import "fmt"

func main() {
	var r, p float32

	fmt.Print("r = ")
	fmt.Scan(&r)

	fmt.Print("p= ")
	fmt.Scan(&p)

	hisob(r, p)
}
func hisob(x float32, y float32) {

	L := 2 * x * y
	S := x * x * y
	fmt.Println("L = ", L, " ", "S = ", S)
}
