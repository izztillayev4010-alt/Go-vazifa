/*
Sana: 21-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Array yoki slice elementlari ichidan eng kattasini topuvchi dastur tuzing
*/

package main

import "fmt"

func main() {
	var (
		n, bigger int
	)

	fmt.Print("Array uzunligini kiriting: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	bigger = arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] > bigger {
			bigger = arr[i]
		}
	}
	fmt.Println("array elementlarining ichida eng kattasi: ", bigger)
}
