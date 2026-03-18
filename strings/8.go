/*
Sana: 17/03/2026
Muallif: Izzatillayev Komronbek
Maqsad: Berilgan gap ichidan polindromlarini topuvchi dastur tuzing
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var gap string = "non shirin"
	res := strings.Split(gap, " ")
	for i := 0; i < len(res); i++ {
		reverse1(res[i])
	}

}
func reverse1(s string) {
	res := make([]byte, len(s))
	left := 0

	for i := len(s) - 1; i >= 0; i-- {
		res[left] = s[i]

		left++
	}
	if s == string(res) {
		fmt.Println("Polindrom")
	} else {
		fmt.Println("polindrom emas")
	}

}
