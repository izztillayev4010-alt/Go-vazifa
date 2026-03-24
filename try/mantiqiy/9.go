/*
Sana: 19-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: A soni berilgan! Jumlani rostlikka tekshiring! A soni 3 xonali va toq son
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

	b = a > 99 && a < 1000

	fmt.Println("a soni 3 xonali: ", b)
}
