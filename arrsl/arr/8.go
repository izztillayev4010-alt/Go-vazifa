// Sana: 13-03-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Array yoki slice elementlari ichidagi musbat va manfiy sonlarni alohida slicelarga yig’uvchi dastur tuzing topuvchi dastur tuzing
//(faqat array bilan)

package main

import "fmt"

func main() {
	var a [5]int = [5]int{}
	var musbat [5]int
	var manfiy [5]int
	for i := 0; i < len(a); i++ {
		fmt.Scan(&a[i])
	}

	for i := 0; i < len(a); i++ {
		if a[i] > 0 {
			musbat[i] = a[i]
		} else if a[i] < 0 {
			manfiy[i] = a[i]
		}

	}
	fmt.Println(musbat, manfiy)
}
