/*
Sana: 19-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Uch xonali a soni berilgan jumlani rostlikka tekshiring: a soni polindrom va juf son
*/

package main

import "fmt"

func main() {
	var (
		a int
		b bool
	)

	fmt.Print("a = ")
	fmt.Scan(&a)

	if a > 99 && a < 1000 {
		b = a/100 == a%10 && a%2 == 0
		fmt.Println("a soni juft va polindrom: ", b)
	} else {
		fmt.Println("a soni 3 xonali emas!")
		return
	}
}
