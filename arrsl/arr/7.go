// Sana: 13-03-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Array yoki slice elementlari ichidagi juft va toq sonlarni alohida slicelarga yig’uvchi dastur tuzing topuvchi dastur tuzing
//(faqat array bilan)

package main

import "fmt"

func main() {
	var a [5]int = [5]int{}
	var juft [5]int
	var toq [5]int
	for i := 0; i < len(a); i++ {
		fmt.Scan(&a[i])
	}

	for i := 0; i < len(a); i++ {
		if a[i]%2 == 0 {
			juft[i] = a[i]
		} else if a[i]%2 != 0 {
			toq[i] = a[i]
		}

	}
	fmt.Println(juft, toq)
}
