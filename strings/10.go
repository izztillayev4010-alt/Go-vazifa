/*
Sana: 17/03/2026
Muallif: Izzatillayev Komronbek
Maqsad: Berilgan gapning ichidan berilgan so’z necha marta takrorlanganligini aniqlovchi dastur tuzing
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var soz string = "salom"
	var gap string = "salom salom salom salom"
	salom(soz, gap)

}
func salom(s, r string) {

	count := 0
	res := strings.Split(r, " ")
	for i := 0; i < len(res); i++ {
		if string(res[i]) == s {
			count++
		} else {
			count = count
		}

	}
	fmt.Println("Takrorlanish soni: ", count)
}
