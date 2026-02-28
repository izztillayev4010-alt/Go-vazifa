// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Son kiritiladi (1, 10) kiriltilgan sonni so’zlar orqali ifodalovchi dastur tuzing
//(Qiymat qaytaradigan func bilan)

package main

import "fmt"

func main() {
	var a int

	fmt.Print("a = ")
	fmt.Scan(&a)

	b := hisob(a)
	fmt.Println(b)
}
func hisob(x int) string {
	switch x {
	case 1:
		return "Bir"
	case 2:
		return "Ikki"
	case 3:
		return "Uch"
	case 4:
		return "To'rt"
	case 5:
		return "Besh"
	case 6:
		return "Olti"
	case 7:
		return "Yetti"
	case 8:
		return "Sakkiz"
	case 9:
		return "To'qqiz"
	case 10:
		return "O'n"
	}
	{
		return "."
	}
}
