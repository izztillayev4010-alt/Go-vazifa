// Sana: 20-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: To’rt xonali a soni berilgan! Jumlani rostilkka tekshiring! A sonining juft o’rindagi raqamlari yigindisi va toq o’rindiadi ramlari yig’indisining ayrimasi 0 ga teng va a soning xar-bir raqami takrorlanmas

package main

import "fmt"

func main() {
	var a int
	
	fmt.Print("a = ")
	fmt.Scan(&a)

	b := a / 1000
	c := a / 100 % 10
	d := a / 10 % 10
	e := a % 10
	f := c + e +( d - b)==0


}