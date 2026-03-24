/*
Sana: 21-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Berilgan so’z necha bo’g’indan iborat ekanligini aniqlovchi dastur tuzing
*/

package main

import "fmt"

func main() {
	var (
		soz   string = "mutolaa"
		count int
	)

	byte := []byte(soz)

	for i := 0; i < len(byte); i++ {
		switch byte[i] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'U':
			count++
		}
	}
	fmt.Println("bo'g'inlar soni: ", count)
}
