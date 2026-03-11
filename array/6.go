//Sana: 11-03-2026
//Muallif: Izzatillayev Komronbek
//Maqsad: Array yoki slice elementlari ichidan juftlarini topuvchi dastur tuzing

package main

import "fmt"

func main() {

	var arr [5]int = [5]int{}

	for i := 0; i < 5; i++ {
		fmt.Scan(&arr[i])
	}
	s := 0

	for i := 0; i < 5; i++ {
		if arr[i]%2 == 0 {
			s = arr[i]
			fmt.Println("juft son: ", s)
		} else {
			fmt.Println("juftson emas: ", arr[i])
		}

	}

}
