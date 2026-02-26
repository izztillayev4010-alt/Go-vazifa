// Sana: 20-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad:  a va b sonlari berilgan! Bu sonlarning yig’indisi musbatligini va toqligini aniqlovchi dastur tuzing!

package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("a = ")
	fmt.Scan(&a)

	fmt.Print("b = ")
	fmt.Scan(&b)

	c := (a+b) > 0 && (a+b)%2 != 0

	fmt.Println("c: ", c)
}
