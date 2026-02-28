// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Berilgan ikki xonali soni raqamlari o’rini almashtirishdan xosil bo’lgan sonni toping!
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
func hisob(x int) int {
	y := x/10 + x%10*10
	return y
}
