/*
sana: 19-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Berilgan 4 xonali sonning raqamlarini teskari tartibda yozish orqali xosil bo’ladigan sonni toping!
*/

package main

import "fmt"

func main() {
	var (
		a, teskari int
	)

	fmt.Print("a = ")
	fmt.Scan(&a)

	if a > 999 && a < 10000 {
		teskari = a/1000 + a/100%10*10 + a/10%10*100 + a%10*1000

		fmt.Println("Teskarisi= ", teskari)
	} else {
		fmt.Println("a soni 4 xonali emas!")

		return
	}
}
