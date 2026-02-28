// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Son kiritiladi (10, 100) kiriltilgan sonni so’zlar orqali ifodalovchi dastur tuzing
//(Qiymat qaytaradigan func bilan)

package main

import "fmt"

func main() {
	var a int

	fmt.Print("a = ")
	fmt.Scan(&a)

	b := a / 10
	c := a % 10
	msg := ""
	d := hisob(b, c, msg)
	fmt.Println(d)
}
func hisob(x int, y int, z string) string {
	switch x {
	case 1:
		z = "O'n"
	case 2:
		z = "Yigirma"
	case 3:
		z = "O'ttiz"
	case 4:
		z = "Qirq"
	case 5:
		z = "Ellik"
	case 6:
		z = "Oltimish"
	case 7:
		z = "Yetmish"
	case 8:
		z = "Sakson"
	case 9:
		z = "To'qson"
	}
	switch y {
	case 1:
		z = z + " bir"
	case 2:
		z = z + " ikki"
	case 3:
		z = z + " uch"
	case 4:
		z = z + " to'rt"
	case 5:
		z = z + " besh"
	case 6:
		z = z + " olti"
	case 7:
		z = z + " yetti"
	case 8:
		z = z + " sakkiz"
	case 9:
		z = z + " to'qqiz"

	}
	return z
}
