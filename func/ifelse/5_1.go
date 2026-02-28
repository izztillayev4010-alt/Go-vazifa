// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: a soni berilgan! Agar a soni 3 ga va 5 ga bo’linsa FIZZBUZZ, agar faqat 3 ga bo’linsa FIZZ, agar faqat 5 ga bo’linssa BUZZ so’zini ekranga chiqaruchi dastur tuzing!
//(Qiymat qaytarmaydigan func bilan)

package main

import "fmt"

func main() {
	var a int

	fmt.Print("a = ")
	fmt.Scan(&a)

	hisob(a)
}

func hisob(x int) {
	if x%3 == 0 && x%5 == 0 {
		fmt.Println("FIZZBUZZ")
	} else if x%3 == 0 {
		fmt.Println("FIZZ")
	} else if x%5 == 0 {
		fmt.Println("BUZZ")
	}
}
