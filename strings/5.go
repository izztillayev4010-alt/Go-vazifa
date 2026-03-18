/*
Sana: 17/03/2026
Muallif: Izzatillayev Komronbek
Maqsad: Berilgan gapni byte arrayga aylantiruvchi dastur tuzing
*/

package main

import (
	"fmt"
)

func main() {
	var gap string = "A"

	byteArray := []byte(gap)
	fmt.Println("Gapning byte arrayi: ", byteArray)
}
