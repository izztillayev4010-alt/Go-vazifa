/*
Sana: 21-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Berilgan gap ichidan polindromlarini topuvchi dastur tuzing
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		gap string = "          bolta arra     "
	)

	sozlar := strings.Split(gap, " ")

	for _, soz := range sozlar {
		if soz != "" {
			if polin(soz) {
				fmt.Println("so'z polindrom!", soz)
			} else {
				fmt.Println("polindrom emas!", soz)
			}
		}
	}

}
func polin(x string) bool {
	soz := []byte(x)

	var (
		oxiri, boshi = len(soz) - 1, 0
	)

	for boshi < oxiri {
		if soz[boshi] != soz[oxiri] {
			return false
		}
		boshi++
		oxiri--
	}

	return true

}
