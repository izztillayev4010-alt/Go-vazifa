/*
Sana: 17/03/2026
Muallif: Izzatillayev Komronbek
Maqsad: Berilgan gap nechta so’zdan iborat ekanligini topuvchi dastur tuzing
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var gap string = "Salom dunyo"

	sozlar := strings.Split(gap, " ")
	fmt.Println("Gap nechta so’zdan iborat: ", len(sozlar))
}
