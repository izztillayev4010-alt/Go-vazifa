// Sana: 27-02-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: a, b va c sonlari berilgan. Bu sonlarning eng kattasini topuvchi dastur tuzing!
//(Qiymat qaytaradigan func bilan)

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

	d := hisob(a, b, c)
	fmt.Println(d)

}
func hisob(x int, y int, z int) int {
	if x > y && y > z || z > y && y < x {
		return x
	} else if y > x && z < x || z > x && y > x {
		return y
	} else {
		return z
	}
}
