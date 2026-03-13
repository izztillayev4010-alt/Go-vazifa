// Sana: 13-03-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Array yoki slice elementlari ichidan juftlarnini topuvchi dastur tuzing
//(faqat array bilan)

package main

import "fmt"

func main() {
	var a [5]int = [5]int{}

	for i := 0; i < len(a); i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < len(a); i++ {
		if a[i]%2 == 0 {
			fmt.Println(a[i])
		}
	}
}
