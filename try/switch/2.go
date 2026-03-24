/*
Sana: 21-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Son kiritiladi (10, 100) kiriltilgan sonni so’zlar orqali ifodalovchi dastur tuzing
*/

package main

import "fmt"

func main() {
	var a int

	fmt.Print("a = ")
	fmt.Scan(&a)

	b := a / 10
	c := a % 10

	if a > 9 && a < 101 {
		switch b {
		case 1:
			fmt.Print("O'n ")
		case 2:
			fmt.Print("Yigirma ")
		case 3:
			fmt.Print("O'ttiz ")
		case 4:
			fmt.Print("Qirq ")
		case 5:
			fmt.Print("Ellik ")
		case 6:
			fmt.Print("Oltimish ")
		case 7:
			fmt.Print("Yetmish ")
		case 8:
			fmt.Print("Sakson ")
		case 9:
			fmt.Print("To'qson ")
		case 10:
			fmt.Print("Yuz")
		}

		switch c {
		case 1:
			fmt.Println("Bir")
		case 2:
			fmt.Println("Ikki")
		case 3:
			fmt.Println("Uch")
		case 4:
			fmt.Println("To'rt")
		case 5:
			fmt.Println("Besh")
		case 6:
			fmt.Println("Olti")
		case 7:
			fmt.Println("Yetti")
		case 8:
			fmt.Println("Sakkiz")
		case 9:
			fmt.Println("To'qqiz")
		}
	} else {
		fmt.Println("Kiritilgan son 10-100 orasida emas!")
	}
}
