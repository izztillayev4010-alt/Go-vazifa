/*
Sana: 21-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Berilgan so’zda qatnashgan unli va undosh harflar sonini aniqlovchi dastur tuzing
*/

package main

import "fmt"

func main() {
	var (
		soz          string = "    salomat"
		unli, undosh int
	)

	byte := []byte(soz)

	for i := 0; i < len(byte); i++ {
		if byte[i] >= 'A' && byte[i] <= 'z' {
			if byte[i] == 'a' || byte[i] == 'e' || byte[i] == 'i' || byte[i] == 'o' || byte[i] == 'u' || byte[i] == 'A' || byte[i] == 'E' || byte[i] == 'I' || byte[i] == 'O' || byte[i] == 'U' {
				unli++
			} else {
				undosh++
			}
		}
	}

	fmt.Println("undoshlar soni: ", undosh, " ", "unlilar soni: ", unli)
}
