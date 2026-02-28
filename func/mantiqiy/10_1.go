// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: To’rt xonali a soni berilgan! Jumlani rostilkka tekshiring! A sonining juft o’rindagi raqamlari yigindisi va toq o’rindiadi ramlari yig’indisining ayrimasi 0 ga teng va a soning xar-bir raqami takrorlanmas
// (Qiymat qaytarmaydigan func bilan)

package main

import "fmt"

func main() {
	var a int

	fmt.Print("a = ")
	fmt.Scan(&a)

	hisob(a)

}
func hisob(x int) {
	m := x / 1000
	n := x / 100 % 10
	o := x / 10 % 10
	p := x % 10
	y := (n+p-m-o) == 0 && (m != n && m != o && m != p && n != 0 && n != p && o != p)
	fmt.Println(y)

}
