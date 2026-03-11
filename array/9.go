//Sana: 11-03-2026
//Muallif: Izzatillayev Komronbek
//Maqsad: Array yoki slice e’lementlari 0 dan boshlab bittadan n gacha sonlardan iborat faqat bitta son tushib qoladi! Tushib qolgan sonni topuvchi dastur tuzing

package main

import "fmt"

func main() {
	var arr [5]int = [5]int{}

	for i := 0; i < 5; i++ {
		fmt.Scan(&arr[i])
	}
	sum := 6 * (1 + arr[4]) / 2
	arrsum := 0

	for i := 0; i < 5; i++ {
		arrsum = arrsum + arr[i]
	}
	res := sum - arrsum

	fmt.Println("Tushib qolgan son: ", res)
}
