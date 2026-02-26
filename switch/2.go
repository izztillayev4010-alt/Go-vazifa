// Sana: 20-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Son kiritiladi (10, 100) kiriltilgan sonni so’zlar orqali ifodalovchi dastur tuzing

package main

import "fmt"

func main() {
	var a int

	fmt.Print("a = ")
	fmt.Scan(&a)

	b := a / 10
	c := a % 10
	msg := ""
	switch b {
	case 1:
		msg = "O'n"
	case 2:
		msg = "Yigirma"
	case 3:
		msg = "O'ttiz"
	case 4:
		msg = "Qirq"
	case 5:
		msg = "Ellik"
	case 6:
		msg = "Oltimish"
	case 7:
		msg = "Yetmish"
	case 8:
		msg = "Sakson"
	case 9:
		msg = "To'qson"
	}
	switch c {
	case 1:
		msg = msg + " bir"
	case 2:
		msg = msg + " ikki"
	case 3:
		msg = msg + " uch"
	case 4:
		msg = msg + " to'rt"
	case 5:
		msg = msg + " besh"
	case 6:
		msg = msg + " olti"
	case 7:
		msg = msg + " yetti"
	case 8:
		msg = msg + " sakkiz"
	case 9:
		msg = msg + " to'qqiz"

	}
	fmt.Println(msg)
}
