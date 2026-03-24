/*
Sana: 21-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Array yoki slice elementlari ichidan toqlarnini topuvchi dastur tuzing
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

	toq := make([]int, n)

	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			toq[i] = arr[i]
		}
	}
	fmt.Println("array elementlari ichida toqlari: ", toq)
}
