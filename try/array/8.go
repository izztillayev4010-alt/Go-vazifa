/*
Sana: 21-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Array yoki slice elementlari ichidagi musbat va manfiy sonlarni alohida
slicelarga yig’uvchi dastur tuzing topuvchi dastur tuzing
*/

package main

import "fmt"

func main() {
	var (
		n int
	)

	fmt.Print("Array uzunligini kiriting: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	musbat := make([]int, n)

	manfiy := make([]int, n)

	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			manfiy[i] = arr[i]
		} else if arr[i] > 0 {
			musbat[i] = arr[i]
		}
	}

	fmt.Println("musbatlari: ", musbat, "manfiylari: ", manfiy)
}
