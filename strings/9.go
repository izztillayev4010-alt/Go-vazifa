/*
Sana: 17/03/2026
Muallif: Izzatillayev Komronbek
Maqsad: Berilgan gapning ichidan berilgan so’z bor yoki yo’qligini topuvchi dastur tuzing
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var soz string = "salom"
	var gap string = "Abduvali bilan Asliddin kokatlar"
	result := salom(soz, gap)

	if result {
		fmt.Println("bor", soz)
	} else {
		fmt.Println("yoq")
	}

}
func salom(s, r string) bool {
	var a bool
	res := strings.Split(r, " ")
	for i := 0; i < len(res); i++ {
		if string(res[i]) == s {
			a = true
		} else {
			a = false
		}

	}
	return a
}
