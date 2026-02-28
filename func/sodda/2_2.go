// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Doiraning radiusi berilgan r, uning aylanasi uzunligi va yuzini hisblovchi dastur tuzing
//(Qiymat qaytaradigan func bilan)

package main

import "fmt"

func main() {
	var r, p float32

	fmt.Print("r = ")
	fmt.Scan(&r)

	fmt.Print("p= ")
	fmt.Scan(&p)

	m := hisob1(r, p)
	n := hisob2(r, p)

	fmt.Println("L = ", m, " ", "S = ", n)
}
func hisob1(x float32, y float32) float32 {

	L := 2 * x * y
	return L

}
func hisob2(x float32, y float32) float32 {
	S := x * x * y
	return S
}
