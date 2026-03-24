/*
Sana: 21-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Berilgan so’zning barcha harflarini agar kichik harflardan iborat
bo’lsa katta harfga aylantiruvchi aks holda kichik harflarga aylantiruvchi dastur tuzing
*/

package main

import "fmt"

func main() {
	var (
		soz string = "   saLOm"
	)

	sozbyte := []byte(soz)

	for i := 0; i < len(sozbyte); i++ {
		if sozbyte[i] >= 32 {
			if sozbyte[i] >= 'a' && sozbyte[i] <= 'z' {
				sozbyte[i] = sozbyte[i] - 32
			} else if sozbyte[i] >= 'A' && sozbyte[i] <= 'Z' {
				sozbyte[i] = sozbyte[i] + 32
			}
		}
	}

	fmt.Printf("%s/n", sozbyte)
}
