// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Berilgan ikki xonali soni raqamlari o’rini almashtirishdan xosil bo’lgan sonni toping!
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
	y := x/10 + x%10*10
	fmt.Println(y)
}
