// Sana: 20-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad:  a va b sonlarni berilgan jumlani rostlikka tekshiring: a soni b soniga teng

package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("a = ")
	fmt.Scan(&a)

	fmt.Print("b = ")
	fmt.Scan(&b)

	c := a == b

	fmt.Println("a soni b soniga tengmi: ", c)
}
