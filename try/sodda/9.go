/*
sana: 19-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Berilgan 4 xonali sonning juft o’rindagi raqamlari yi’gindisidan toq o’rindagi raqamlari yig’indising ayrimasini topuvchi dastur tuzing
*/

package main

import "fmt"

func main() {
	var (
		a, ayirma int
	)

	fmt.Print("a = ")
	fmt.Scan(&a)

	if a > 999 && a < 10000 {
		ayirma = a/100%10 + a%10 - a/1000 - a/10%10

		fmt.Println("Ayirma: ", ayirma)
	} else {
		fmt.Println("a soni 4 xonali emas")
		return
	}
}
