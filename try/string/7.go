/*
Sana: 21-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Kiritilgan so’zni polindrom yoki polindrom emasligini aniqlovchi dastur tuzing
*/

package main

import "fmt"

func main() {
	var (
		soz string = "bolta"
	)

	if ispolindrom(soz) {
		fmt.Println("so'z polindrom! ", soz)
	} else {
		fmt.Println("polindrom emas!")
	}

}
func ispolindrom(x string) bool {
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
