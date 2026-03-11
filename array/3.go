//Sana: 11-03-2026
//Muallif: Izzatillayev Komronbek
//Maqsad: Array yoki slice elementlari ichidan eng kattasini topuvchi dastur tuzing

package main

import "fmt"

func main() {

	var arr [5]int = [5]int{}

	for i := 0; i < 5; i++ {
		fmt.Scan(&arr[i])
	}
	s := arr[0]

	for i := 0; i < 5; i++ {
		for j := 1; j < 5; j++ {
			if arr[i] < arr[j] {
				s = arr[j]
			}
		}
	}
	fmt.Println(s)
}
