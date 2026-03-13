// Sana: 13-03-2026
// Muallif: Izzatillayev Komronbek
// Maqsad:Berilgan 4 xonali sonning juft o’rindagi raqamlari yi’gindisidan toq o’rindagi raqamlari yig’indising ayrimasini topuvchi dastur tuzing
//(pointer bilan)

package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("a = ")
	fmt.Scan(&a)

	hisob(&a, &b)

	fmt.Println(b)
}
func hisob(x, y *int) {
	*y = *x/100%10 + *x%10 - *x/1000 - *x/10%10

}
