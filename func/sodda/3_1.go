// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: L uzunlik berilgan sm larda! Shu uzunlik ichidagi metrlar sonini topuvchi dastur tuzing!
//(Qiymat qaytarmaydigan func bilan)

package main

import "fmt"

func main() {
	var l int

	fmt.Print("l = ")
	fmt.Scan(&l)

	hisob(l)

}
func hisob(x int) {
	L := x / 100
	fmt.Println("L = ", L)
}
