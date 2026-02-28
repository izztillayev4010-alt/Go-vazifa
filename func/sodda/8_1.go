// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Berilgan uch xonali sonni onlar va yuzlar xonasidagi raqamlari o’rnini alishtirishdan xosil bo’lgan sonni toping
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
	y := x/100*10 + x/10%10*100 + x%10
	fmt.Println(y)
}
