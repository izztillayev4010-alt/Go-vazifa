//Sana: 11-03-2026
//Muallif: Izzatillayev Komronbek
//Maqsad: Array yoki slice elementlarini teskari tartibda dastur tuzing

package main

import "fmt"

func main() {

	var arr [5]int = [5]int{}
	var arr1 [5]int = [5]int{}

	for i := 0; i < 5; i++ {
		fmt.Scan(&arr[i])
	}
	arr1 = arr
	s := 4
	for i := 0; i < 5; i++ {

		arr[i] = arr1[s]

		s--
	}
	fmt.Println(arr)
}
