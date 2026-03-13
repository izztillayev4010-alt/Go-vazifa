// Sana: 13-03-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Berilgan ikki xonali soni raqamlari o’rini almashtirishdan xosil bo’lgan sonni toping!
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
func hisob(m, n *int) {
	*n = *m%10*10 + *m/10

}
