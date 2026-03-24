/*
Sana: 19-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: To’rt xonali a soni berilgan! Jumlani rostilkka tekshiring! A sonining juft o’rindagi raqamlari yigindisi va toq o’rindiadi ramlari yig’indisining ayrimasi 0 ga teng va a soning xar-bir raqami takrorlanmas
*/

package main

import "fmt"

func main() {
	var (
		a int
		b bool
	)

	fmt.Print("a = ")
	fmt.Scan(&a)

	if a > 999 && a < 10000 {
		var arr []int = []int{a / 1000, a / 100 % 10, a / 10 % 10, a % 10}

		b = (a/100%10 + a%10 - a/1000 - a/10%10) == 0

		for i := 0; i < len(arr); i++ {
			for j := i + 1; j < len(arr); j++ {

				if arr[i] == arr[j] {
					b = false
				} else {
					b = b
				}
			}

		}
		fmt.Println("a soning juft o'rinda turganlardan toq o'rinda turganlarining ayirmasi 0 va raqamlari takrorlanmas: ", b)
	} else {
		fmt.Println("a soni 4 xonali emas!")
	}

}
