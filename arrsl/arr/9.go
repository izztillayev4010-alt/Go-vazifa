// Sana: 13-03-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Array yoki slice e’lementlari 0 dan boshlab bittadan n gacha sonlardan iborat faqat bitta son tushib qoladi! Tushib qolgan sonni topuvchi dastur tuzing
//(faqat array bilan)

package main

import "fmt"

func main() {
	var a [5]int = [5]int{}

	for i := 0; i < len(a); i++ {
		fmt.Scan(&a[i])
	}
	sum := a[4] * (a[4] + a[0]) / 2
	sumarr := 0
	for i := 0; i < len(a); i++ {
		sumarr += a[i]
	}
	res := sum - sumarr
	fmt.Println("tushib qolgani", res)
}
