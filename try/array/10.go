/*
Sana: 21-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Array yoki slice elementlari uchida polindromlarini
alohida slicega yig’uvchi dastur tuzing!
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

	polindrom := make([]int, n)

	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < len(arr); i++ {
		if ispolindrom(arr[i]) {
			if arr[i] >= 0 && arr[i] < 10 {
				continue
			}
			polindrom[i] = arr[i]
		}
	}
	fmt.Println("array elementlari orasida polindromlari: ", polindrom)
}
func ispolindrom(x int) bool {
	r := 0
	reverse := 0
	org := x

	for x > 0 {
		r = x % 10

		reverse = reverse*10 + r

		x = x / 10
	}
	if reverse == org {
		return true
	} else {
		return false
	}
}
