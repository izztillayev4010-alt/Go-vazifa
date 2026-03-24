/*
Sana: 21-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Array yoki slice elementlari yig’indisni topuvchi dastur tuzing
*/

package main

import "fmt"

func main() {
	var (
		n, summa int
	)

	fmt.Print("Array uzunligini kiritng: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < len(arr); i++ {
		summa += arr[i]
	}

	fmt.Println("Array elementlari yig'indisi: ", summa)
}
