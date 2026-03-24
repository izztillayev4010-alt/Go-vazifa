/*
sana: 19-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Berilgan ikki xonali sonni raqamlari yig’indisini toping!
*/

package main

import "fmt"

func main() {

	var (
		a, sum int
	)

	fmt.Print("a = ")
	fmt.Scan(&a)

	if a > 9 && a < 100 {

		sum = a/10 + a%10

		fmt.Println("Yig'indi: ", sum)
	} else {
		fmt.Println("a soni 2 xonali son emas!")
		return
	}
}
