/*
Sana: 17/03/2026
Muallif: Izzatillayev Komronbek
Maqsad: berilgan so’zni birinchi harfini katta harfga o’giruvchi dastur tuzing
*/

package main

import (
	"fmt"
)

func main() {
	var soz string = "salom"

	s := soz[0]
	s = s - 32

	soz = string(s) + soz[1:]

	fmt.Println(" Natija: ", soz)
}
