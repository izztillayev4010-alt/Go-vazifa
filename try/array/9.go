/*
Sana: 21-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Array yoki slice e’lementlari 0 dan boshlab bittadan n gacha
sonlardan iborat faqat bitta son tushib qoladi! Tushib qolgan sonni topuvchi dastur tuzing
*/

package main

import "fmt"

func main() {
	var (
		n, summa, arraysum int
	)

	fmt.Print("Array uzunligini kiriting: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	summa = (n + 2) * (n + 1) / 2

	for i := 0; i < len(arr); i++ {
		arraysum += arr[i]
	}

	summa = summa - arraysum

	fmt.Println("tushib qolgan son: ", summa)
}
