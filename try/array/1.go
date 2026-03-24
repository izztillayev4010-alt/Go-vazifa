/*
Sana: 21-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Array yoki slice elementlarini teskari tartibda dastur tuzing
*/

package main

import "fmt"

func main() {
	var (
		n, m int
	)

	fmt.Print("array uzunligini kiriting: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	reverse := make([]int, n)

	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	for i := len(arr) - 1; i >= 0; i-- {
		reverse[m] = arr[i]

		m++
	}
	fmt.Println(reverse)
}
