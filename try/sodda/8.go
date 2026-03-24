/*
sana: 19-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Berilgan uch xonali sonni onlar va yuzlar xonasidagi raqamlari o’rnini alishtirishdan xosil bo’lgan sonni toping
*/

package main

import "fmt"

func main() {
	var (
		a, ex int
	)

	fmt.Print("a = ")
	fmt.Scan(&a)

	if a > 99 && a < 1000 {
		ex = a/100*10 + a/10%10*100 + a%10

		fmt.Println("Alamashgani: ", ex)
	} else {
		fmt.Println("a soni 3 xonali emas!")
		return
	}
}
