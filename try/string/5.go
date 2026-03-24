/*
Sana: 21-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Berilgan gapni byte arrayga aylantiruvchi dastur tuzing
*/

package main

import "fmt"

func main() {
	var (
		gap string = "salom dunyo"
	)

	gapbyte := []byte(gap)

	fmt.Println(gapbyte)
}
