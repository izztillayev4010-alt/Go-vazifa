// Sana: 20-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Uch xonali a soni berilgan jumlani rostlikka tekshiring: a soni polindrom va juf son

package main

import "fmt"

func main() {
	var a int

	fmt.Print("a = ")
	fmt.Scan(&a)

	b := (a/100 == a%10) && (a%2 == 0)

	fmt.Println("a soni polindrom va juf son: ", b)
}
