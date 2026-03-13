// Sana: 13-03-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Array yoki slice elementlarini teskari tartibda dastur tuzing
//(faqat array bilan)

package main

import "fmt"

func main() {
	var a [5]int = [5]int{}

	for i := 0; i < len(a); i++ {
		fmt.Scan(&a[i])
	}
	for i := len(a) - 1; i >= 0; i-- {
		fmt.Println(a[i])
	}

}
