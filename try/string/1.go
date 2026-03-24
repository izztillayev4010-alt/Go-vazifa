/*
Sana: 21-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: berilgan so’zni birinchi harfini katta harfga o’giruvchi dastur tuzing
*/

package main

import "fmt"

func main() {
	var (
		soz string = "   salom"
	)

	byte := []byte(soz)

	for i := 0; i < len(byte); i++ {
		if byte[i] <= 32 {
			continue
		} else {
			byte[i] = byte[i] - 32
			break
		}

	}

	fmt.Printf("%s\n", byte)
}
