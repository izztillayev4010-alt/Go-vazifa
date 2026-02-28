// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: a va b sonlari berilgan. Bu sonlarning kattasini topuvchi dastur tuzing!
//(Qiymat qaytarmaydigan func bilan)

package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("a = ")
	fmt.Scan(&a)

	fmt.Print("b = ")
	fmt.Scan(&b)
	hisob(a, b)
}
func hisob(x int, y int) {
	if x > y {
		fmt.Println(x)
	} else if y > x {
		fmt.Println(y)
	} else {
		fmt.Println(x, "=", y)
	}

}
