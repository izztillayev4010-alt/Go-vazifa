// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Son kiritiladi (1, 10) kiriltilgan sonni so’zlar orqali ifodalovchi dastur tuzing
//(Qiymat qaytarmaydigan func bilan)

package main

import "fmt"

func main() {
	var a int

	fmt.Print("a = ")
	fmt.Scan(&a)

	hisob(a)

}
func hisob(x int) {
	switch x {
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
	case 10:
		fmt.Println("O'n")
	}
}
