// Sana: 17-92-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: a, b va c sonlari berilgan. Bu sonlarning o’rachasini topuvchi dastur tuzing!

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

	if a > b && b >= c {
		fmt.Println("o'rtachasi ", b)
	} else if b > a && a >= c {
		fmt.Println("o'rtachasi  ", a)
	} else if c > b && b >= a {
		fmt.Println("o'rtachasi  ", b)
	} else if a >= b && b > c {
		fmt.Println("o'rtachasi ", b)
	} else if a >= c && c > b {
		fmt.Println("o'rtachasi ", c)
	} else if c >= a && a > b {
		fmt.Println("o'rtachasi ", a)
	} else {
		fmt.Println("Sonlar o'zaro teng: ", a, " ", b, " ", c)
	}
}
