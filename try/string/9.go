/*
Sana: 21-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Berilgan gapning ichidan berilgan so’z bor yoki yo’qligini topuvchi dastur tuzing
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		gap, soz = "         salom dunyo", "   salom   "
	)

	sozlar := strings.Split(gap, " ")

	words := strings.Split(soz, " ")

	for _, m := range words {
		if m != "" {
			soz = m
		}
	}

	for _, word := range sozlar {
		if word != "" {
			if word == soz {
				fmt.Println("bor", soz)
			}
		}
	}

}
