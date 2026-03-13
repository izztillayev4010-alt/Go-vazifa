// Sana: 13-03-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: L uzunlik berilgan sm larda! Shu uzunlik ichidagi metrlar sonini topuvchi dastur tuzing!
//(pointer bilan)

package main

import "fmt"

func main() {
	var L, res float32

	fmt.Print("L = ")
	fmt.Scan(&L)

	hisob(&L, &res)

	fmt.Println("l = ", res)
}
func hisob(m, n *float32) {
	*n = *m / 100
}
