// Sana: 20-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: A soni berilgan! Jumlani rostlikka tekshiring! A soni 3 xonali va toq son

package main

import "fmt"

func main() {
	var a int

	fmt.Print("a = ")
	fmt.Scan(&a)

	b := (a > 99) && (a < 1000) && a%2 != 0

	fmt.Println("A soni 3 xonali va toq son: ", b)
}
