// Sana: 20-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Berilgan a soni musbatligini aniqlovchi dastur tuzing

package main

import "fmt"

func main() {
	var a int

	fmt.Print("a = ")
	fmt.Scan(&a)

	b := a > 0

	fmt.Println("a soni musbatmi: ", b)
}
