/*
Sana: 17/03/2026
Muallif: Izzatillayev Komronbek
Maqsad: Berilgan so’zning barcha harflarini agar kichik harflardan iborat bo’lsa katta harfga aylantiruvchi aks holda kichik harflarga aylantiruvchi dastur tuzing
*/

package main

import (
	"fmt"
)

func main() {
	var soz string = "saLom"

	result := ""
	for i := 0; i < len(soz); i++ {
		if soz[i] >= 'a' && soz[i] <= 'z' {
			result += string(soz[i] - 32)
		} else if soz[i] >= 'A' && soz[i] <= 'Z' {
			result += string(soz[i] + 32)
		} else {
			result += string(soz[i])
		}
	}

	fmt.Println("Natija: ", result)
}
