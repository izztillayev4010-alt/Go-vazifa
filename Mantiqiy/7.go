// Sana: 20-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad:  Ikki xonal a soni berilgan! Jumlani rostlikka tekshiring: a soni polindrom va toq son

package main

import "fmt"

func main() {
	var a int

	fmt.Print("a = ")
	fmt.Scan(&a)

	b := a / 10
	c := a % 10
	d := (b == c) && a%2 != 0

	fmt.Println("d = ", d)
}
