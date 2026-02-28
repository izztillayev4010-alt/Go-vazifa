// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: L uzunlik berilgan sm larda! Shu uzunlik ichidagi metrlar sonini topuvchi dastur tuzing!
//(Qiymat qaytaradigan func bilan)

package main

import "fmt"

func main() {
	var l int

	fmt.Print("l = ")
	fmt.Scan(&l)

	m := hisob(l)

	fmt.Println("L = ", m)

}
func hisob(x int) int {
	L := x / 100
	return L
}
