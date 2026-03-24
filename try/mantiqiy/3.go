/*
Sana: 19-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: a va b sonlarni berilgan jumlani rostlikka tekshiring: a soni b sonidan katta
*/

package main

import "fmt"

func main() {
	var (
		a, b int
		c    bool
	)

	fmt.Print("a = ")
	fmt.Scan(&a)

	fmt.Print("b = ")
	fmt.Scan(&b)

	c = a > b

	fmt.Println("a soni b sonidan kattami: ", c)
}
