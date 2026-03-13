// Sana: 13-03-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Berilgan uch xonali sonni onlar va yuzlar xonasidagi raqamlari o’rnini alishtirishdan xosil bo’lgan sonni toping
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
	*y = *x/100*10 + *x/10%10*100 + *x%10

}
