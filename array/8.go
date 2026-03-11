//Sana: 11-03-2026
//Muallif: Izzatillayev Komronbek
//Maqsad: Array yoki slice elementlari ichidagi musbat va manfiy sonlarni alohida slicelarga yig’uvchi dastur tuzing topuvchi dastur tuzing

package main

import "fmt"

func main() {

	var arr [5]int = [5]int{}

	for i := 0; i < 5; i++ {
		fmt.Scan(&arr[i])
	}
	var musbat [5]int = [5]int{}
	var manfiy [5]int = [5]int{}

	for i := 0; i < 5; i++ {
		if arr[i] > 0 {
			musbat[i] = arr[i]

		}

	}
	for i := 0; i < 5; i++ {
		if arr[i] < 0 {
			manfiy[i] = arr[i]

		}

	}
	fmt.Println("musbat sonlar: ", musbat)
	fmt.Println("manfiy sonlar: ", manfiy)
}
