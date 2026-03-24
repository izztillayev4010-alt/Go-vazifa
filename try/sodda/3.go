/*
sana: 19-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: L uzunlik berilgan sm larda! Shu uzunlik ichidagi metrlar sonini topuvchi dastur tuzing!
*/

package main

import "fmt"

func main() {
	var l float32

	fmt.Print("uzunlikni kiriting(sm da): ")
	fmt.Scan(&l)

	L := l / 100

	fmt.Println("L = ", L, " metr")
}
