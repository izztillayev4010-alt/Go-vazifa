/*
Sana: 19-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Ikki xonal a soni berilgan! Jumlani rostlikka tekshiring: a soni polindrom va toq son
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
	if a > 9 && a < 100 {
		b = a/10 == a%10 && a%2 != 0
		fmt.Println("a soni toq va polindrom:", b)
	} else {
		fmt.Println("a soni 2 xonali emas!")
		return
	}

}
