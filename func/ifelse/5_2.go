// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: a soni berilgan! Agar a soni 3 ga va 5 ga bo’linsa FIZZBUZZ, agar faqat 3 ga bo’linsa FIZZ, agar faqat 5 ga bo’linssa BUZZ so’zini ekranga chiqaruchi dastur tuzing!
//(Qiymat qaytaradigan func bilan)

package main

import "fmt"

func main() {
	var a int

	fmt.Print("a = ")
	fmt.Scan(&a)

	n := hisob(a)
	fmt.Println(n)
}

func hisob(x int) string {
	if x%3 == 0 && x%5 == 0 {
		m := "FIZZBUZZ"
		return m
	} else if x%3 == 0 {
		o := "FIZZ"
		return o
	} else {
		d := "BUZZ"
		return d
	}
}
