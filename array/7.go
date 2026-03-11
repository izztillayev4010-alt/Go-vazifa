//Sana: 11-03-2026
//Muallif: Izzatillayev Komronbek
//Maqsad: Array yoki slice elementlari ichidagi juft va toq sonlarni alohida slicelarga yig’uvchi dastur tuzing topuvchi dastur tuzing

package main

import "fmt"

func main() {

	var arr [5]int = [5]int{}

	for i := 0; i < 5; i++ {
		fmt.Scan(&arr[i])
	}
	var juft [5]int = [5]int{}
	var toq [5]int = [5]int{}

	for i := 0; i < 5; i++ {
		if arr[i]%2 == 0 {
			juft[i] = arr[i]

		}

	}
	for i := 0; i < 5; i++ {
		if arr[i]%2 != 0 {
			toq[i] = arr[i]

		}

	}
	fmt.Println("juft sonlar: ", juft)
	fmt.Println("toq sonlar: ", toq)
}
