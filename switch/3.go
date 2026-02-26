// Sana: 20-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Son kiritiladi (1, 1000) kiriltilgan sonni so’zlar orqali ifodalovchi dastur tuzing

package main

import "fmt"

func main() {
	var a int

	fmt.Print("a = ")
	fmt.Scan(&a)

	b := a / 100
	c := a / 10 % 10
	d := a % 10
	msg := ""
	switch b {
	case 1:
		msg = "Bir yuz"
	case 2:
		msg = "Ikki yuzi"
	case 3:
		msg = "Uch yuzi"
	case 4:
		msg = "To'rt yuzi"
	case 5:
		msg = "Besh yuzi"
	case 6:
		msg = "Olti yuzi"
	case 7:
		msg = "Yetti yuzi"
	case 8:
		msg = "Sakkiz yuzi"
	case 9:
		msg = "To'qqiz yuzi"
	}

	switch c {
	case 1:
		msg = msg + " O'n"
	case 2:
		msg = msg + " Yigirma"
	case 3:
		msg = msg + " O'ttiz "
	case 4:
		msg = msg + " Qirq "
	case 5:
		msg = msg + " Ellik "
	case 6:
		msg = msg + " Oltimish "
	case 7:
		msg = msg + " Yettimish "
	case 8:
		msg = msg + " Sakson "
	case 9:
		msg = msg + " To'qson "
	}
	switch d {
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
