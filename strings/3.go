/*
Sana: 17/03/2026
Muallif: Izzatillayev Komronbek
Maqsad: Berilgan so’zda qatnashgan unli va undosh harflar sonini aniqlovchi dastur tuzing
*/

package main

import (
	"fmt"
)

func main() {
	var soz string = "komronbek"
	unli := 0
	undosh := 0
	for i := 0; i < len(soz); i++ {
		if soz[i] == 'a' || soz[i] == 'e' || soz[i] == 'i' || soz[i] == 'o' || soz[i] == 'u' {
			unli++
		} else {
			undosh++
		}
	}
	fmt.Println("Unli harflar soni: ", unli)
	fmt.Println("Undosh harflar soni: ", undosh)

}
