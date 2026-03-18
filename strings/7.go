/*
Sana: 17/03/2026
Muallif: Izzatillayev Komronbek
Maqsad: Kiritilgan so’zni polindrom yoki polindrom emasligini aniqlovchi dastur tuzing
*/

package main

import "fmt"

func main() {
	var soz string = "Abduvali"

	result := reverse(soz)

	if result {
		fmt.Println("Polindrom", soz)
	} else {
		fmt.Println("polindrom emas", soz)
	}
}
func reverse(s string) bool {
	res := make([]byte, len(s))
	left := 0

	for i := len(s) - 1; i >= 0; i-- {
		res[left] = s[i]

		left++
	}
	if s == string(res) {
		return true
	}
	return false
}
