/*
Sana: 21-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Array yoki slice elementlari ichidan eng kichigini topuvchi dastur tuzing
*/

package main

import "fmt"

func main() {
	var (
		n, smaller int
	)

	fmt.Print("Array uzunligini kiriting: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	smaller = arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] < smaller {
			smaller = arr[i]
		}
	}
	fmt.Println("array elementlarining ichida eng kichigi: ", smaller)
}
