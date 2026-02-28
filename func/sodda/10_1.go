// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Berilgan 4 xonali sonning raqamlarini teskari tartibda yozish orqali xosil bo’ladigan sonni toping!
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
	y := x/100%10*10 + x%10*1000 + x/1000 + x/10%10*100
	fmt.Println(y)
}
