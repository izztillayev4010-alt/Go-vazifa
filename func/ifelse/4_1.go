// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: a, b va c sonlari berilgan. Bu sonlarning eng kichigini topuvchi dastur tuzing!
//(Qiymat qaytarmaydigan func bilan)

package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Print("a = ")
	fmt.Scan(&a)

	fmt.Print("b = ")
	fmt.Scan(&b)

	fmt.Print("c = ")
	fmt.Scan(&c)

	hisob(a, b, c)

}
func hisob(x int, y int, z int) {
	if x > y && y > z || z < x && y > x {
		fmt.Println(z)
	} else if y < x && z < x || z > x && y < x {
		fmt.Println(y)
	} else {
		fmt.Println(x)
	}
}
