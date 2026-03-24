/*
Sana: 21-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Berilgan gap nechta so’zdan iborat ekanligini topuvchi dastur tuzing
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		gap   string = "     salom dunyo"
		count int
	)

	sozlar := strings.Split(gap, " ")

	for _, soz := range sozlar {
		if soz != "" {
			count++
		}
	}

	fmt.Println("sozlar soni: ", count)
}
