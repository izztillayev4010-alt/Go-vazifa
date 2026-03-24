/*
sana: 19-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Berilgan ikki xonali soni raqamlari o’rini almashtirishdan xosil bo’lgan sonni toping!
*/

package main

import "fmt"

func main() {

	var (
		a, ex int
	)

	fmt.Print("a = ")
	fmt.Scan(&a)

	if a > 9 && a < 100 {

		ex = a/10 + a%10*10

		fmt.Println("Almashgani: ", ex)
	} else {
		fmt.Println("a soni 2 xonali son emas!")
		return
	}
}
