// Sana: 13-03-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Array yoki slice elementlari ichidan eng kattasini topuvchi dastur tuzing
//(faqat array bilan)

package main

import "fmt"

func main() {
	var a [5]int = [5]int{}

	for i := 0; i < len(a); i++ {
		fmt.Scan(&a[i])
	}
	s := a[0]
	for i := 0; i < len(a); i++ {
		for j := 1; j < len(a); j++ {
			if a[i] < a[j] {
				s = a[j]
			}
		}
	}
	fmt.Println(s)
}
