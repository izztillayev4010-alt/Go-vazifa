/*
Sana: 20-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: a soni berilgan! Agar a soni 3 ga va 5 ga bo’linsa FIZZBUZZ, agar faqat 3 ga
bo’linsa FIZZ, agar faqat 5 ga bo’linssa BUZZ so’zini ekranga chiqaruchi dastur tuzing!
*/

package main

import "fmt"

func main() {
	var a int

	fmt.Print("a = ")
	fmt.Scan(&a)

	if a%15 == 0 {
		fmt.Println("FIZZBUZZ! ")
	} else if a%3 == 0 {
		fmt.Println("FIZZ")
	} else if a%5 == 0 {
		fmt.Println("BUZZ")

	} else {
		fmt.Println("3ga ham 5ga ham bo'linmaydi!")
	}
}
