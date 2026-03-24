/*
Sana: 21-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Ikkilik sanoq tizimida berilgan sonni o’nlik sanoq tizimiga o’tkazuchi dastur tuzing
*/

package main

import "fmt"

func main() {
	var ikkilik int

	fmt.Print("Ikkillik son kiriitng: ")
	fmt.Scan(&ikkilik)

	res := twototen(ikkilik)

	fmt.Println(res)

}
func twototen(x int) int {
	r := 0

	s := 1
	result := 0
	for x > 0 {
		r = x % 10
		result += s * r
		s = s * 2
		x = x / 10
	}
	return result
}
