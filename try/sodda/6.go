/*
sana: 19-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Berilgan uch xonali sonni raqamlari yig’ndisni toping!
*/

package main

import "fmt"

func main() {

	var (
		a, sum int
	)

	fmt.Print("a = ")
	fmt.Scan(&a)

	if a > 99 && a < 1000 {

		sum = a/100 + a/10%10 + a%10

		fmt.Println("Yig'indi: ", sum)
	} else {
		fmt.Println("a soni 3 xonali son emas!")
		return
	}
}
