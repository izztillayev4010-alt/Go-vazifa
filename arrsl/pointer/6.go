// Sana: 13-03-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Berilgan uch xonali sonni raqamlari yig’ndisni toping!
//(pointer bilan)

package main

import "fmt"

func main() {
	var a, res int

	fmt.Print("a = ")
	fmt.Scan(&a)

	hisob(&a, &res)

	fmt.Println("Yig'indi: ", res)

}
func hisob(b, r *int) {
	*r = *b/100 + *b/10%10 + *b%10
}
