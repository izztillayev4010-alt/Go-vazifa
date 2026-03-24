/*
Sana: 21-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: o’nlik sanoq tizimida berilgan sonni ikkilik sanoq tizimiga o’tkazuchi dastur tuzing
*/

package main

import "fmt"

func main() {
	var onlik int

	fmt.Print("O'nlikdagi son kiriting: ")
	fmt.Scan(&onlik)

	result := tentotwo(onlik)

	fmt.Println(result)
}
func tentotwo(x int) string {
	res := ""
	for i := 0; i < x; i++ {
		if x%2 == 0 {
			res = "0" + res
		} else if x%2 == 1 {
			res = "1" + res
		}
		x = (x - x%2) / 2
		if x/2 < 1 && x == 1 {
			res = "1" + res
		} else if x/2 < 1 && x == 0 {
			res = "0" + res
		}
	}
	count := 0
	for i := 0; i < len(res); i++ {
		if res[i] == '0' {
			count++
		}
	}
	if count == len(res) {
		res = "1" + res + "0"
	}
	return string(x) + res
}
