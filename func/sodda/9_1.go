// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Berilgan 4 xonali sonning juft o’rindagi raqamlari yi’gindisidan toq o’rindagi raqamlari yig’indising ayrimasini topuvchi dastur tuzing
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
	y := x/100%10 + x%10 - x/1000 - x/10%10
	fmt.Println(y)
}
