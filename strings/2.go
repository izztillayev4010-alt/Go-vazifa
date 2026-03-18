/*
Sana: 17/03/2026
Muallif: Izzatillayev Komronbek
Maqsad: Berilgan so’z necha bo’g’indan iborat ekanligini aniqlovchi dastur tuzing
*/

package main

import (
	"fmt"
)

func main() {
	var soz string = "mutolaa"

	count := 0
	for i := 0; i < len(soz); i++ {
		if soz[i] == 'a' || soz[i] == 'e' || soz[i] == 'i' || soz[i] == 'o' || soz[i] == 'u' {
			count++
		}
	}

	fmt.Println("So’z necha bo’g’indan iborat: ", count)
}
